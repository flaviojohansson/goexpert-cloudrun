# goexpert-cloudrun
Desafio Fullcycle - Pós GoExpert - Labs - Deploy com Cloud Run

## Testar o projeto
```

git clone https://github.com/flaviojohansson/goexpert-cloudrun

cd goexpert-cloudrun

# Criar arquivo .env com a chave de API de weatherapi.com
echo 'WEATHER_API_KEY="your-api-key"' > .env

go run .

curl http://localhost:8080/temperatura?cep=80530000
```

## Testes automatizados
```
# uma vez configurado o arquivo .env com a chave de API de weatherapi.com:
go test
```

## Testar com Docker compose
- Primeiro, adicionar a chave de API de weatherapi.com no arquivo docker-compose.yaml
```
docker compose up -d
curl http://localhost:8080/temperatura?cep=80530000
```

## Testar ao vivo no Google Cloud Run
- A chave de API deste live example expira em 06/04/2025

```
curl https://temperatura-service-651650408590.southamerica-west1.run.app/temperatura?cep=80530000
```
