# nsm-api

api untuk mendapatkan Nomor Statistik Madrasah (NSM) seluruh Indonesia


# NSM
* NSM Object
```
{
  nsm: string,
  name: string,
  city: string,
  district: string
}
```


### Get NSM List  
`GET /nsms`  

Returns a paginated list of NSM objects.  

#### Query Parameters  
| Param  | Type   | Default | Description                          |  
|--------|--------|---------|--------------------------------------|  
| `page` | `int`  | `1`     | Page number (1-based).               |  
| `limit`| `int`  | `10`    | Items per page (max: 100).           |  

#### Request Headers  
```http
Content-Type: application/json  
Accept: application/json  
{
  "success": true,
  "data": [
    { "id": 1, "name": "NSM 1", ... },
    { "id": 2, "name": "NSM 2", ... }
  ],
  "pagination": {
    "total": 50,
    "page": 1,
    "limit": 10,
    "total_pages": 5
  }
}
```

