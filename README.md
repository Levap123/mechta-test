 # MECHTA test

Тестовое задание на позицию golang backend developer в компанию <strong>Mechta.kz</strong>

# Usage

```sh
go run . <файл с данными> <количество воркеров>
```

## Example
```sh
go run . data.txt 3
```
## Как сгенерировать файл data.txt

```sh
chmod +x scripts/gen_data.sh
scripts/gen_data.sh 2000
```
Где число - количество объектов
