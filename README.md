# marketplace-ordersvc

POST /orders
req:  { "userId": "<uuid>", "items": [ { "sku": "SKU1", "qty": 1 } ] }
201:  { "id": "<uuid>", "status": "NEW", "userId": "<uuid>", "items": [...], "createdAt": "<rfc3339>" }
409:  { "error": "duplicate" }  # если пришёл тот же Idempotency-Key

GET /orders/{id}
200:  { ...Order }
404:  { "error": "not_found" }
Headers: Idempotency-Key (опц.) — для POST