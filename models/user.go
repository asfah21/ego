package models

// User merepresentasikan data pengguna yang mengisi kuesioner
type User struct {
	ID                string `json:"id"`
	Nama              string `json:"nama"`
	Email             string `json:"email"`
	SkorNarsisme      int    `json:"skor_narsisme"`
	SkorMachiavellian int    `json:"skor_machiavellian"`
	SkorPsikopati     int    `json:"skor_psikopati"`
	StatusPembayaran  string `json:"status_pembayaran"`
}

// QuizResult adalah data yang dikirim ke template hasil
type QuizResult struct {
	Nama          string
	Narsisme      int
	Machiavellian int
	Psikopati     int
}

// PaywallData adalah data yang dikirim ke template paywall
type PaywallData struct {
	ID   string
	Nama string
}
