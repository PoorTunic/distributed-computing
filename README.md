# distributed-computing

1. Téléchargez le serveur de base de données redis:
    http://download.redis.io/releases/redis-5.0.3.tar.gz
2. Après décompresser le fichier
3. Dans le dossier redis utiliser les commandes
    make test (pour faire l'instalation)
    src/redis-server
4. Ouvrez une terminal et lancez la commande dans le dossier central, cela lancera le serveur principal
    go run central.go
5. Ouvres une autre terminal et lancez la commande dans le dossier project, cela lancera un client
    go run client-server.go -port <# de Port>

6. Les chemins pour tester les serveurs sont
    GET: http://localhost:8080/ || Cela obtient tous les enregistrements
    GET: http://localhost:8080/{id} || Cela obtient l'information d'un enregistrement selon son ID
    PUT: http://localhost:8080/{id} || JSON fields : "data" || Cela ajoute l'enregistrement avec l'ID dans le URL et le corps donné ou cela met à jour l'enregistrement si cela déjà existe

NOTE: Tous les URLs sont égal pour le client, il faut changer le nombre de port

