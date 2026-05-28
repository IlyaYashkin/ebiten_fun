# Интерактивное демо физики частиц

Интерактивное демо физики частиц на Go и [Ebitengine](https://ebitengine.org/). Частицы следуют за курсором с простой физикой движения.

## Демо

https://github.com/user-attachments/assets/379a514a-ca9c-4fcc-a32f-00d97804e1f1

## Как это работает

Частицы разбиты на батчи, каждый из которых обрабатывается в отдельной горутине. Финальный кадр отрисовывается за один вызов при помощи WritePixels, что обеспечивает высокую производительность даже при большом количестве частиц.

## Стек

- [Go 1.23](https://go.dev/)
- [Ebitengine v2](https://ebitengine.org/)

## Запуск

**Требования:** Go 1.23+

```bash
git clone https://github.com/IlyaYashkin/ebiten_fun.git
cd ebiten_fun
go run main.go
```

## Лицензия

MIT
