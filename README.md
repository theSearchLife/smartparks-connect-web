# smartparks-connect-web
Web application for remotely managing Smart Parks OpenCollar sensors through Chirpstack and Rockblock.

## QuickStart
1. Compile front-end project： ```cd public && npm install && npm run build && cd ../```
2. Start server ```go run main.go```

## Development

### Backend
To start a development server for the backend you need to run ```go run main.go```

### Frontend
To work on the front end the best way is to run the development version of the front end.
```cd public && npm run dev```. Frontend will be running on http://localhost:5173/

To build frontend for production run ```cd public && npm install && npm run build```
