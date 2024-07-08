package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "e98908f6d9151f1ec970085c3d914cf9"

// GeocodeResponse represents the response structure for geocode API
type GeocodeResponse struct {
	Status   string `json:"status"`
	Geocodes []struct {
		Location string `json:"location"`
	} `json:"geocodes"`
}

// GetGeocode returns the latitude and longitude for a given address
func GetGeocode(address string) (string, error) {
	url := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?address=%s&key=%s", address, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var geocodeResponse GeocodeResponse
	err = json.Unmarshal(body, &geocodeResponse)
	if err != nil {
		return "", err
	}
	if geocodeResponse.Status == "1" && len(geocodeResponse.Geocodes) > 0 {
		//fmt.Println(fmt.Printf("%+v", geocodeResponse))
		return geocodeResponse.Geocodes[0].Location, nil
	}
	return "", fmt.Errorf("failed to get geocode")
}

// DrivingRouteResponse represents the response structure for driving route API
type DrivingRouteResponse struct {
	Status string `json:"status"`
	Route  struct {
		Paths []struct {
			Distance string `json:"distance"`
			Duration string `json:"duration"`
			Tolls    string `json:"tolls"`
		} `json:"paths"`
		TaxiCost string `json:"taxi_cost"`
	} `json:"route"`
}

// GetDrivingRoute returns the distance and duration between two locations with optional waypoints
func GetDrivingRoute(origin, destination string, waypoints ...string) (string, string, string, string, error) {
	var waypointsParam string
	if len(waypoints) > 0 {
		waypointsParam = fmt.Sprintf("&waypoints=%s", waypoints[0])
		for _, waypoint := range waypoints[1:] {
			waypointsParam += fmt.Sprintf(";%s", waypoint)
		}
	}

	url := fmt.Sprintf("https://restapi.amap.com/v3/direction/driving?origin=%s&destination=%s%s&extensions=all&key=%s", origin, destination, waypointsParam, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", err
	}

	var routeResponse DrivingRouteResponse
	err = json.Unmarshal(body, &routeResponse)
	if err != nil {
		return "", "", "", "", err
	}
	if routeResponse.Status == "1" && len(routeResponse.Route.Paths) > 0 {
		//fmt.Printf("%+v", routeResponse)
		return routeResponse.Route.Paths[0].Distance, routeResponse.Route.Paths[0].Duration, routeResponse.Route.TaxiCost, routeResponse.Route.Paths[0].Tolls, nil
	}
	return "", "", "", "", fmt.Errorf("failed to get driving route")
}

func main() {
	startAddress := "天津市神州浩天科技有限公司"
	endAddress := "天津市南站火车站"
	waypointsAddresses := []string{"天津市滨海新区核心商务区", "天津市滨江道商业街"}

	// 获取起点位置
	startLocation, err := GetGeocode(startAddress)
	if err != nil {
		fmt.Println("Error getting start location:", err)
		return
	}

	// 获取终点位置
	endLocation, err := GetGeocode(endAddress)
	if err != nil {
		fmt.Println("Error getting end location:", err)
		return
	}

	// 获取途径点位置
	var waypointsLocations []string
	for _, address := range waypointsAddresses {
		location, err := GetGeocode(address)
		if err != nil {
			fmt.Printf("Error getting waypoint location for %s: %v\n", address, err)
			return
		}
		waypointsLocations = append(waypointsLocations, location)
	}

	// 获取路径信息
	distance, duration, taxiCost, tolls, err := GetDrivingRoute(startLocation, endLocation, waypointsLocations...)
	if err != nil {
		fmt.Println("Error getting driving route:", err)
		return
	}

	//fmt.Printf("从 %s 到 %s 的行车距离为 %s 米，预计耗时 %v 分钟。(预计打车费用 %s 元, 道路收费 %s 元)\n", startAddress, endAddress, distance, t.New(duration).Int32()/60, taxiCost, tolls)

	// 打印途径点信息
	fmt.Println("行车起点:", startAddress)
	fmt.Println("途径点:")
	for i, address := range waypointsAddresses {
		fmt.Printf("%d. %s \n", i+1, address)
	}
	fmt.Println("行车终点:", endAddress)
	// 打印行车路径信息
	fmt.Printf("行车距离为 %s 米，预计耗时 %v 分钟。(预计打车费用 %s 元, 道路收费 %s 元)\n", distance, duration, taxiCost, tolls)
}
