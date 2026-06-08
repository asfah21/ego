package services

import (
	"ego/models"
	"ego/repositories"
)

// ProcessQuizAnswers memproses jawaban kuesioner dan menyimpan ke database
func ProcessQuizAnswers(email, nama string, q1, q2, q3, q4, q5 int) (string, error) {
	// Konversi skor 1-4 ke persentase 25-100
	// Q1, Q4 -> Narsisme
	// Q2, Q5 -> Machiavellian
	// Q3 -> Psikopati
	skorNarsisme := (q1 + q4) * 25 / 2
	skorMachiavellian := (q2 + q5) * 25 / 2
	skorPsikopati := q3 * 25

	userID, err := repositories.InsertUser(email, nama, skorNarsisme, skorMachiavellian, skorPsikopati)
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

// ConfirmPayment mengonfirmasi pembayaran user
func ConfirmPayment(id string) error {
	return repositories.UpdatePaymentStatus(id)
}
