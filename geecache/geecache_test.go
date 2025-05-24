package geecache

import (
	"log"
	"reflect"
	"testing"
	"fmt"
)

var db = map[string]string{
	"Tom": "630",
	"Jeck": "589",
	"Sam": "560",
}

func TestGetter(t *testing.T){

	var f Getter = GetterFunc(func(key string)([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")

	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect){
		t.Fatal("callback failed")
	}
}

func TestGet(t *testing.T){
	loadCounts := make(map[string]int, len(db))

	gee := NewGroup("scores", 2<<10, GetterFunc(func(key string)([]byte, error){
		log.Println("Load data:", key)
		if v, ok := db[key]; ok {
			if _, ok := loadCounts[key]; !ok {
				loadCounts[key] = 0
			}
			loadCounts[key] ++
			return []byte(v) , nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get value from Tom")
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatal("cache not work")
		} 
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknown should be empty, but %s got", view)
	}
}

func TestGroup(t *testing.T){
	groupName := "scores"
	NewGroup(groupName, 2<<10, GetterFunc(
		func(key string)([]byte, error){return nil, fmt.Errorf("%s not exist", key)}))

	if group := GetGroup(groupName); group == nil  || group.name != groupName {
		t.Fatal("failed to get group")
	}

	if group := GetGroup(groupName + "111"); group != nil {
		t.Fatalf("expected nil, but %s got",group.name)
	}
}