package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVillagesByName(t *testing.T) {

	req, err := http.NewRequest("GET", "/villages", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "latiung")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnVillages)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"data":[{"id":"1101010001","name":"LATIUNG","district":"TEUPAH SELATAN","regency":"KABUPATEN SIMEULUE","province":"ACEH"}],"meta":{"total_items":1,"items_per_page":20,"current_page":1,"total_page":1}}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetVillagesByNameWithPageAndPageSize(t *testing.T) {

	req, err := http.NewRequest("GET", "/villages", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "kebun")
	q.Add("page", "4")
	q.Add("page_size", "27")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnVillages)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"data":[{"id":"1223010004","name":"PERKEBUNAN BERANGIR","district":"NA IX-X","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223020001","name":"PERKEBUNAN PERNANTIAN","district":"MARBAU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223020002","name":"PERKEBUNAN MARBAU SELATAN","district":"MARBAU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223020003","name":"PERKEBUNAN MILANO","district":"MARBAU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223020016","name":"PERKEBUNAN BRUSSEL","district":"MARBAU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223030005","name":"PERKEBUNAN PADANG HALABAN","district":"AEK KUO","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223040006","name":"PERKEBUNAN AEK PAMINGKE","district":"AEK NATAS","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223040008","name":"PERKEBUNAN HALIMBE","district":"AEK NATAS","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223050006","name":"PERKEBUNAN DAMULI","district":"KUALUH SELATAN","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223070003","name":"PERKEBUNAN LONDUT","district":"KUALUH HULU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223070004","name":"PERKEBUNAN KANOPAN ULU","district":"KUALUH HULU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223070007","name":"PERKEBUNAN MEMBANG MUDA","district":"KUALUH HULU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223070008","name":"PERKEBUNAN LABUHAN HAJI","district":"KUALUH HULU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1223070009","name":"PERKEBUNAN HANNA","district":"KUALUH HULU","regency":"KABUPATEN LABUHAN BATU UTARA","province":"SUMATERA UTARA"},{"id":"1273050001","name":"KEBUN SAYUR","district":"SIANTAR TIMUR","regency":"KOTA PEMATANG SIANTAR","province":"SUMATERA UTARA"},{"id":"1276040006","name":"KEBUN LADA","district":"BINJAI UTARA","regency":"KOTA BINJAI","province":"SUMATERA UTARA"},{"id":"1277010017","name":"PERKEBUNAN P KOLING","district":"PADANGSIDIMPUAN TENGGARA","regency":"KOTA PADANGSIDIMPUAN","province":"SUMATERA UTARA"},{"id":"1401020011","name":"KEBUN LADO","district":"SINGINGI","regency":"KABUPATEN KUANTAN SINGINGI","province":"RIAU"},{"id":"1401052007","name":"LUBUK KEBUN","district":"LOGAS TANAH DARAT","regency":"KABUPATEN KUANTAN SINGINGI","province":"RIAU"},{"id":"1402042002","name":"PERKEBUNAN  SUNGAI LALA","district":"SUNGAI LALA","regency":"KABUPATEN INDRAGIRI HULU","province":"RIAU"},{"id":"1402042007","name":"PERKEBUNAN SUNGAI PARIT","district":"SUNGAI LALA","regency":"KABUPATEN INDRAGIRI HULU","province":"RIAU"},{"id":"1406011010","name":"KEBUN TINGGI","district":"KAMPAR KIRI HULU","regency":"KABUPATEN KAMPAR","province":"RIAU"},{"id":"1406013001","name":"KEBUN DURIAN","district":"GUNUNG SAHILAN","regency":"KABUPATEN KAMPAR","province":"RIAU"},{"id":"1409011034","name":"PERKEBUNAN SIARANGARANG","district":"PUJUD","regency":"KABUPATEN ROKAN HILIR","province":"RIAU"},{"id":"1409014010","name":"PERKEBUNAN TANJUNG MEDAN","district":"TANJUNG MEDAN","regency":"KABUPATEN ROKAN HILIR","province":"RIAU"},{"id":"1501010009","name":"KEBUN BARU","district":"GUNUNG RAYA","regency":"KABUPATEN KERINCI","province":"JAMBI"},{"id":"1501010018","name":"KEBUN LIMA","district":"GUNUNG RAYA","regency":"KABUPATEN KERINCI","province":"JAMBI"}],"meta":{"total_items":137,"items_per_page":27,"current_page":4,"total_page":6}}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetVillagesByNameButNoData(t *testing.T) {

	req, err := http.NewRequest("GET", "/villages", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "newyork")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnVillages)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"data":null,"meta":{"total_items":0,"items_per_page":20,"current_page":1,"total_page":0}}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
