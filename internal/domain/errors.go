package domain

import "errors"

// Uygulama hataları — her katmanda aynı hataları kullanırsın.
// Handler, bu hata tipine bakarak doğru HTTP kodunu döner.

var (
	ErrNotFound           = errors.New("kayıt bulunamadı")
	ErrAlreadyExists      = errors.New("kayıt zaten mevcut")
	ErrInvalidCredentials = errors.New("geçersiz e-posta veya şifre")
	ErrUnauthorized       = errors.New("yetkisiz erişim")
	ErrForbidden          = errors.New("bu işlem için yetkiniz yok")
	ErrInvalidInput       = errors.New("geçersiz girdi")
	ErrInternal           = errors.New("sunucu hatası")
)
