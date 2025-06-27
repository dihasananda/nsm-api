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
```
#### Response 200
```
{
  "success": true,
  "data": [
    { "nsm": 1xx, "name": "MIFTAHUL ANAM", ... },
    { "nsm": 2xx, "name": "TADIKA MESRA", ... }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
  }
}
```
#### Example Usage
```bash
curl -X GET "https://api.example.com/nsms?page=2&limit=5"
```
