package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func PrintUrl(queryMap map[string]string, path string, msg string) string{
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	//var u net.URL

	u.Scheme = "http"
	u.Host = "127.0.0.1:8080"
	u.Path = path
	q := u.Query()
	for key, val := range queryMap {
		q.Set(key, val)
	}

	u.RawQuery = q.Encode()
	fmt.Println(msg, u)
	return u.String()

}

//variant url1:  http://127.0.0.1:8080/variant?discount_price=100&mrp=549.2222&product_id=22&q=dotnet&size=30ml&variant_id=36&variant_name=quantity
//variant url2:  http://127.0.0.1:8080/variant?mrp=459.2222&product_id=22&q=dotnet&size=60ml&variant_id=39&variant_name=quantity
//variant url3:  http://127.0.0.1:8080/variant?color=black&discount_price=345.6&mrp=949.2222&product_id=26&q=dotnet&size=large&variant_id=40
//variant url4:  http://127.0.0.1:8080/variant?color=black&mrp=249.2222&product_id=26&q=dotnet&size=medium&variant_id=50
//variant url5:  http://127.0.0.1:8080/variant?mrp=767&product_id=15&q=dotnet&size=90ml&variant_id=59
//product url1:  http://127.0.0.1:8080/product?category_ids=3&description=frient+for+your+hair&product_id=22&product_name=shampoo&q=dotnet&url=https%3A%2F%2Fwww.jabing.in%2Fshampo-%3Fdc
//product url2:  http://127.0.0.1:8080/product?category_ids=3&description=smothing+for+hair&product_id=26&product_name=conditioner&q=dotnet&url=https%3A%2F%2Fwww.myntra.in%2Fcmooth-%3Fdchild%3D1%26keywords%3Dpuma%2Bshirt%26qid%3D1610287417%26sr%3D8-3
//product url3:  http://127.0.0.1:8080/product?category_ids=3%2C1&description=grow+hair+again&product_id=15&product_name=hair+growth&q=dotnet&url=https%3A%2F%2Fwww.amazon.in%2Fgrowth-%3Fdchild%3D1%26keywords%3Dpuma%2Bshirt%26qid%3D1610287417%26sr%3D8-3
//category url1:  http://127.0.0.1:8080/category?category_id=0&category_name=amazon&q=dotnet
//category url2:  http://127.0.0.1:8080/category?category_id=1&category_name=cosmetic&parent_category_id=0&q=dotnet
//category url3:  http://127.0.0.1:8080/category?category_id=2&category_name=education&parent_category_id=0&q=dotnet
//category url4:  http://127.0.0.1:8080/category?category_id=3&category_name=hair&parent_category_id=1&q=dotnet

//variant update url1:  http://127.0.0.1:8080/variant/36?mrp=90&product_id=22&q=dotnet&size=88ml&variant_name=quantity




func GenerateUrl(){
	urlList := []string{}
	var_path := "variant"
	pro_path:= "product"
	cat_path := "category"

	var_1 := map[string]string{
		"variant_id": "36",
		"variant_name": "quantity",
		"mrp": "549.2222",
		"size": "30ml",
		"discount_price": "100",
		"product_id": "22",
	}
	var_2 := map[string]string{
		"variant_id": "39",
		"variant_name": "quantity",
		"mrp": "459.2222",
		"size": "60ml",
		"product_id": "22",
	}
	var_3 := map[string]string{
		"variant_id": "40",
		"mrp": "949.2222",
		"size": "large",
		"color": "black",
		"discount_price": "345.6",
		"product_id": "26",
	}
	var_4 := map[string]string{
		"variant_id": "50",
		"mrp": "249.2222",
		"size": "medium",
		"color": "black",
		"product_id": "26",
	}

	var_5 := map[string]string{
		"variant_id": "59",
		"mrp": "767",
		"size": "90ml",
		"product_id": "15",
	}

	urlList = append(urlList, PrintUrl(var_1, var_path,"variant url1: "))
	urlList = append(urlList,PrintUrl(var_2, var_path,"variant url2: "))
	urlList = append(urlList,PrintUrl(var_3, var_path,"variant url3: "))
	urlList = append(urlList,PrintUrl(var_4, var_path,"variant url4: "))
	urlList = append(urlList,PrintUrl(var_5, var_path,"variant url5: "))





	prod_1 := map[string]string{
		"product_id": "22",
		"product_name": "shampoo",
		"description": "frient for your hair",
		"url": "https://www.jabing.in/shampo-?dc",
		"category_ids": "3",
	}
	prod_2 := map[string]string{
		"product_id": "26",
		"product_name": "conditioner",
		"description": "smothing for hair",
		"url": "https://www.myntra.in/cmooth-?dchild=1&keywords=puma+shirt&qid=1610287417&sr=8-3",
		"category_ids": "3",
	}

	prod_3 := map[string]string{
		"product_id": "15",
		"product_name": "hair growth",
		"description": "grow hair again",
		"url": "https://www.amazon.in/growth-?dchild=1&keywords=puma+shirt&qid=1610287417&sr=8-3",
		"category_ids": "3,1",
	}


	urlList = append(urlList,PrintUrl(prod_1, pro_path,"product url1: "))
	urlList = append(urlList,PrintUrl(prod_2, pro_path,"product url2: "))
	urlList = append(urlList,PrintUrl(prod_3, pro_path,"product url3: "))


	cat_1 := map[string]string{
		"category_id": "0",
		"category_name": "amazon",
	}

	cat_2 := map[string]string{
		"category_id": "1",
		"category_name": "cosmetic",
		"parent_category_id": "0",
	}

	cat_3 := map[string]string{
		"category_id": "2",
		"category_name": "education",
		"parent_category_id": "0",
	}
	cat_4 := map[string]string{
		"category_id": "3",
		"category_name": "hair",
		"parent_category_id": "1",
	}


	urlList = append(urlList,PrintUrl(cat_1, cat_path,"category url1: "))
	urlList = append(urlList,PrintUrl(cat_2, cat_path,"category url2: "))
	urlList = append(urlList,PrintUrl(cat_3, cat_path,"category url3: "))
	urlList = append(urlList,PrintUrl(cat_4, cat_path,"category url4: "))

	//sendPostReq(urlList)

	var_update_path := "variant/36"

	//var_1 := map[string]string{
	//	"variant_id": "36",
	//	"variant_name": "quantity",
	//	"mrp": "549.2222",
	//	"size": "30ml",
	//	"discount_price": "100",
	//	"product_id": "22",
	//}
	var_update_1 := map[string]string{
		"variant_name": "quantity",
		"mrp": "90",
		"size": "88ml",
		"product_id": "22",
	}
	urlList = append(urlList,PrintUrl(var_update_1, var_update_path,"variant update url1: "))

}

func main(){
	GenerateUrl()

}

func sendPostReq(urlList []string) {
	for _, url := range urlList {
		log.Println("===================================")
		log.Println("Url: ", url)
		//Encode the data
		postBody, _ := json.Marshal(map[string]string{

		})
		responseBody := bytes.NewBuffer(postBody)
		//Leverage Go's HTTP Post function to make request
		resp, err := http.Post(url, "application/json", responseBody)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	}


}