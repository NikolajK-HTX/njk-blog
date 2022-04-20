---
title: "Kvidrer - app"
date: 2022-04-20T17:03:00Z
draft: false
---

Pakker, der ønskes installeret (via ``dnf`` eller ``apt``):

```bash
git dotnet nginx nano postgresql-server
```

# Opsætning af bruger

Man starter med at logge ind som root på serveren. Så kan man sætte brugeren op:

```bash
$ adduser manager
$ passwd manager
$ usermod -aG wheel manager
$ rsync --archive --chown=manager:manager ~/.ssh /home/manager
```

Man kan skifte til den nye bruger med:

```bash
$ su - manager
```

# Opsætning af database
Der bruges postgresql-server som database (husk at installere det før). Mere information findes på [https://www.postgresql.org/docs/current/creating-cluster.html](postgresql.org)

```bash
$ su - postgres
$ initdb -D /usr/local/pgsql/data
```

# Opsætning af nginx

Nginx blev installeret før. Det kører ikke automatisk, så kør:
```bash
$ sudo systemctl enable --now nginx
```

Man finder konfigurationsfilen på stien \verb+/etc/nginx/nginx.conf+.
```
$ sudo nano /etc/nginx/nginx.conf
```

Hvis nginx bruges som reverse proxy, så prøver man formentlig også at få nginx til at tilgå porte ud over 80 og 443. Det er som udgangspunkt ikke tilladt af SELinux. Derfor skal man tillade det med:

```bash
$ setsebool -P httpd_can_network_connect 1
```

# Dotnet API

```bash
$ dotnet build --configuration Release
$ dotnet run [filepath]/[name].dll
```

# HTTPS via certbot

For at installere certbot på Rocky Linux skal epel-release være installeret. Den giver adgang til flere pakker administreret af brugerne/fællesskabet.

```bash
$ sudo dnf install epel-release
```

Dernæst installeres certbot:
```bash
$ sudo dnf install certbot python3-certbot-nginx
```

Certbot køres med kommandoerne:
```bash
$ sudo certbot --nginx
$ sudo certbot renew --dry-run
```
