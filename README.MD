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

** GET /nsms**
----
 Returns all users in the system.
* **URL Params**  
  None
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Success Response:**  
* **Code:** 200  
  **Content:**  
```
{
  nsms: [
           {<nsm_object>},
           {<nsm_object>},
           {<nsm_object>}
         ]
}
```
