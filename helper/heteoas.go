package helper

import (
	"fmt"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

func CreateLinksForItem(id int, resouse string) []web.Link {
	return []web.Link{
		{Rel: "self", Href: fmt.Sprintf("/%s/%d", resouse, id)},
		{Rel: "update", Href: fmt.Sprintf("/%s/%d", resouse, id)},
		{Rel: "delete", Href: fmt.Sprintf("/%s/%d", resouse, id)},
	}
}

func CreateLinksForItems(ids []int, entity string) []web.Link {
	var allLinks []web.Link
	for _, id := range ids {
		allLinks = append(allLinks, CreateLinksForItem(id, entity)...)
	}
	return allLinks
}
