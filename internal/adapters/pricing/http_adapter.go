package pricing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/madrabit/marketplace-ordersvc/internal/orders"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	hc      *http.Client
	baseUrl string
}

func NewPricingController(baseUrl string) *Controller {
	return &Controller{
		baseUrl: baseUrl,
		hc: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (p *Controller) PriceBatch(ctx context.Context, items []orders.ItemQty) map[uuid.UUID]orders.PriceItem {
	ids := make([]string, 0, len(items))
	for _, i := range items {
		ids = append(ids, i.Id.String())
	}
	fmt.Println(ids)
	body, err := json.Marshal(struct {
		Ids []string `json:"ids"`
	}{
		Ids: ids,
	})
	if err != nil {
		return nil
	}
	log.Println("JSON body:", string(body))
	url := p.baseUrl + "/api/v1/catalogs"
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	fmt.Println(httpReq.URL)
	if err != nil {
		fmt.Printf("ERROR creating request: %v\n", err)
		return nil
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := p.hc.Do(httpReq)
	if err != nil {
		log.Println("write my error")
		log.Printf("DEBUG: HTTP request error: %v\n", err)
		log.Printf("DEBUG: Error type: %T\n", err)
		return nil
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()

	out := make(map[uuid.UUID]orders.PriceItem, len(items))
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return nil
	}
	fmt.Println(out)
	return out
}
