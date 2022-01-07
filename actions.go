package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/urfave/cli/v2"
	"github.com/xmapst/NexusImageClean/nexus"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	wg     *WaitGroup
	client nexus.Client
	err    error
)

func listImagesAction(c *cli.Context) error {
	var nexusVersion = c.String("nexus_version")
	client, err = newConnNexus()
	if err != nil {
		return err
	}
	res, err := client.GetRepository(nexusVersion)
	if err != nil {
		return err
	}
	if len(res.Result.Data) == 0 {

	}

	tpl := table.NewWriter()
	tpl.SetOutputMirror(os.Stdout)
	tpl.SetStyle(table.StyleLight)
	tpl.AppendHeader(table.Row{"Path", "Name", "Type", "Leaf"})
	for _, image := range res.Result.Data {
		tpl.AppendRow([]interface{}{image.ID, image.Text, image.Type, image.Leaf})
	}
	//tpl.AppendSeparator()
	tpl.AppendFooter(table.Row{"Total", len(res.Result.Data), "", ""})
	tpl.Render()
	return nil
}

func listTagsByImage(c *cli.Context) error {
	var imgName = c.String("name")
	var nexusVersion = c.String("nexus_version")
	var nodeID = fmt.Sprintf("%s/%s", nexusVersion, imgName)
	client, err = newConnNexus()
	if err != nil {
		return err
	}
	tagsRes, err := client.GetRepositoryAllTags(nodeID)
	if err != nil {
		return err
	}

	tpl := table.NewWriter()
	tpl.SetOutputMirror(os.Stdout)
	tpl.SetStyle(table.StyleLight)
	tpl.AppendHeader(table.Row{"ComponentID", "AssetID", "Path", "ImageName", "TagName", "Type", "Leaf"})
	for _, tag := range tagsRes.Result.Data {
		tpl.AppendRow([]interface{}{tag.ComponentID, tag.AssetID, tag.ID, imgName, tag.Text, tag.Type, tag.Leaf})
	}
	tpl.AppendFooter(table.Row{"Total", len(tagsRes.Result.Data), "", ""})
	tpl.Render()
	return nil
}

func showImageInfo(c *cli.Context) error {
	var imgName = c.String("name")
	var nexusVersion = c.String("nexus_version")
	var nodeID = fmt.Sprintf("%s/%s", nexusVersion, imgName)
	var tagName = c.String("tag")
	client, err = newConnNexus()
	if err != nil {
		return err
	}
	tagsRes, err := client.GetRepositoryAllTags(nodeID)
	if err != nil {
		return err
	}
	var assetID string
	for _, tag := range tagsRes.Result.Data {
		if tag.Text == tagName {
			assetID = tag.AssetID
			break
		}
	}
	infoRes, err := client.GetRepositoryTagInfo(assetID)
	if err != nil {
		return err
	}

	info, err := yaml.Marshal(&infoRes.Result.Data)
	if err != nil {
		return err
	}
	fmt.Println(string(info))
	return nil
}

func deleteAction(c *cli.Context) error {
	var imgName = c.String("name")
	var nexusVersion = c.String("nexus_version")
	var tagName = c.String("tag")
	var keep = c.Int("keep")
	var current = c.Int("current")

	client, err = newConnNexus()
	if err != nil {
		log.Fatal(err)
	}

	reviewsDataChan := make(chan []ReviewsData)
	defer close(reviewsDataChan)
	if imgName == "" || tagName == "" {
		go func() {
			for {
				data, ok := <-reviewsDataChan
				if !ok {
					return
				}
				for _, tag := range data[:len(data)-keep] {
					nameSlice := strings.Split(tag.Name, "/")
					err := client.DeleteImageTag(tag.ID)
					if err == nil {
						log.Printf("%s image %s tag deleted\n", nameSlice[1], nameSlice[len(nameSlice)-1])
						continue
					}
					log.Println("delete tag error:", err)
				}
			}
		}()
	}

	// 不指定镜像名将全仓进行清理
	if imgName == "" {
		if keep <= 3 {
			return fmt.Errorf("You should either specify the tag or how many images you want to keep")
		}
		// 并发控制
		wg = NewPool(current)
		// 获取仓库下的所有镜像名称
		repo, err := client.GetRepository(nexusVersion)
		if err != nil {
			log.Fatal(err)
		}
		for _, image := range repo.Result.Data {
			// 获取镜像的所有tag
			tags, err := client.GetRepositoryAllTags(image.ID)
			if err != nil {
				if err.Error() != nexus.NotFount {
					log.Println(err)
					log.Printf("%s image failed to get tag\n", image.Text)
				}
				continue
			}
			// 跳过小于数量的仓库
			if len(tags.Result.Data) <= keep {
				continue
			}
			wg.Add(1)
			go GetImageTagList(tags, wg, reviewsDataChan)
		}
		wg.Wait()
		return nil
	}

	var nodeID = fmt.Sprintf("%s/%s", nexusVersion, imgName)
	tagsRes, err := client.GetRepositoryAllTags(nodeID)
	if err != nil {
		return err
	}
	if tagName == "" {
		if keep <= 3 {
			return fmt.Errorf("You should either specify the tag or how many images you want to keep")
		}
		GetImageTagList(tagsRes, nil, reviewsDataChan)
		return nil
	} else {
		var assetID string
		for _, tag := range tagsRes.Result.Data {
			if tag.Text == tagName {
				assetID = tag.AssetID
				break
			}
		}
		if err := client.DeleteImageTag(assetID); err != nil {
			log.Println("delete tag error:", err)
		}
		log.Printf("%s image %s tag deleted\n", imgName, tagName)
	}

	return nil
}

func GetImageTagList(tags *nexus.RepositoryTagsResponse, wg *WaitGroup, reviewsDataChan chan []ReviewsData) {
	if wg != nil {
		defer wg.Done()
	}
	var reviewsData = make(timeSlice, 0, len(tags.Result.Data))
	for _, tag := range tags.Result.Data {
		tagInfo, err := client.GetRepositoryTagInfo(tag.AssetID)
		if err != nil {
			log.Println(err)
			log.Printf("%s image failed to get tag info\n", tag.ID)
			continue
		}
		reviewsData = append(reviewsData, ReviewsData{
			Date:           tagInfo.Result.Data.BlobUpdated.Local(),
			ID:             tagInfo.Result.Data.ID,
			Name:           tagInfo.Result.Data.Name,
			repositoryName: tagInfo.Result.Data.RepositoryName,
		})
	}
	sort.Sort(reviewsData)
	reviewsDataChan <- reviewsData
}
