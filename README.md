### Usage

To Run with live reload and a postgres database image
```
docker compose build && docker compose up
```

Example seed data

```
{
  "title": "Beautiful Apartment",
  "description": "A beautiful apartment in the heart of the city.",
  "rent_price": 2100.50,
  "building_type": "Condo",
  "availability": true,
  "location": {
    "city": "Toronto",
    "street": "5th Avenue",
    "municipality": "North York",
    "unit_no": "21B",
    "postal_code": "10001"
  },
  "utility": {
    "water": true,
    "gas": true,
    "electric": true,
    "parking": 1,
    "locker": true,
    "internet": true,
    "furnished": true
  },
  "features": {
    "rooms": 3,
    "washrooms": 2,
    "kitchen": true,
    "square_feet": 850.5
  }
}

```
