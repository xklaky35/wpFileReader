package filereader

import (
	"os"
	"encoding/json"
)


type Config struct {
	MaxValue int `json:"max_value"`
	MinValue int `json:"min_value"`
	IncreaseStep int `json:"increase_step"`
	DecreaseStep int `json:"decrease_step"`
	Timezone string `json:"timezone"`
}

type Data struct {
	Gauges []Gauge `json:"gauges"`
}

type Gauge struct {
	Name string `json:"name"`
	Value int `json:"value"`
	LastIncrease string `json:"last_increase"`
}


func LoadData(path string) (Data, error){
	var data Data  
	f, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}
		
	err = json.Unmarshal(f, &data)

	return data, nil
}

func LoadConfig(path string) (Config, error){
	var config Config 
	f, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
		
	err = json.Unmarshal(f, &config)

	return config, nil
}

func WriteData(d *Data, path string) error {

	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	
	err = os.WriteFile(path, data, 0666)
	if err != nil{
		return err
	}
	return nil
}

