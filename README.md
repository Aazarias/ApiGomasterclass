
# ApiGomasterclass

The ApiGomasterclass api is used to retrieve main user data (first name, email, nickname) as well as data
concerning his video games, his progress, his library, his wishlist.


## API Reference

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


## Authors

- [@Aazarias](https://github.com/Aazarias)
- Damasio Matthieu

