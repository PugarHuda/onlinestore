// dalam paket database/db.go
package database

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
    // Konfigurasi URI koneksi MongoDB Atlas
    uri := "mongodb+srv://development:testpassword@ecommerce.k1llnha.mongodb.net/?retryWrites=true&w=majority"

    // Buat opsi untuk koneksi MongoDB
    opts := options.Client().ApplyURI(uri)

    // Buat klien baru dan hubungkan ke server
    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        log.Fatal(err)
    }

    // Tentukan konteks dengan batas waktu
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Ping server untuk memastikan koneksi berhasil
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("failed to connect to mongodb:", err)
    }
    fmt.Println("Successfully connected to MongoDB")
    return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
    var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
    return collection
}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection {
    var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
    return productcollection
}
