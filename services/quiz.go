package services

import (
	"ego/models"
	"ego/repositories"
)

// ProcessQuizAnswers memproses jawaban kuesioner dan menyimpan ke database
func ProcessQuizAnswers(nama, email string, q1, q2, q3 int) (string, error) {
	// Konversi skor 1-5 ke persentase 20-100
	skorNarsisme := q1 * 20
	skorMachiavellian := q2 * 20
	skorPsikopati := q3 * 20

	userID, err := repositories.InsertUser(nama, email, skorNarsisme, skorMachiavellian, skorPsikopati)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// GetPaywallData mengambil data untuk halaman paywall
func GetPaywallData(id string) (*models.PaywallData, error) {
	nama, err := repositories.GetUserName(id)
	if err != nil {
		return nil, err
	}
	return &models.PaywallData{ID: id, Nama: nama}, nil
}

// GetQuizResult mengambil data hasil kuis (dengan proteksi paywall)
func GetQuizResult(id string) (*models.QuizResult, error) {
	user, err := repositories.GetUserResult(id)
	if err != nil {
		return nil, err
	}

	// Proteksi: hanya tampilkan hasil jika sudah PAID
	if user.StatusPembayaran != "PAID" {
		return nil, nil // nil menandakan belum bayar
	}

	return &models.QuizResult{
		Nama:          user.Nama,
		Narsisme:      user.SkorNarsisme,
		Machiavellian: user.SkorMachiavellian,
		Psikopati:     user.SkorPsikopati,
	}, nil
}
