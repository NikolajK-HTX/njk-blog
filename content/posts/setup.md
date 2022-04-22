---
title: "Opsætning af blog"
date: 2022-04-10T18:39:35Z
draft: false
---

Denne post er til for at hjælpe med at sætte sådant et projekt op - nok mest for fremtidige mig.

Jeg går ud fra, at man bruger Rocky Linux eller tilsvarende (det kunne f.eks. AlmaLinux, CentOS eller Fedora Server). Det vil dog ikke se voldsomt anderledes ud, hvis man valgte at sætte det op på en anden distro.

Du læser højst sandsynligt dette fra [blog.jehaj.dk](https://blog.jehaj.dk/), som bliver leveret til dig fra en server hos DigitalOcean. Det burde ikke være anderledes at gøre det hos andre tjenester, men hvis man ikke har bestemt sig for en endnu, kan DO anbefales. Jeg har i hvert fald ikke haft nogen problemer med dem.

# Sæt igang
Man kan ikke gå igang før, man har fået adgang til din server. Det gøres med `ssh root@<ip-address>`. I stedet for `<ip-address>` kan man bruge dit domæne, hvis man har sat det til at pege på serveren. DigitalOcean anbefaler man bruger SSH passphrases, så det anbefaler jeg også.

(Det er ikke sikkert, at maskinen er helt opdateret... for en sikkerheds skyld køres `dnf upgrade`)

Dernæst laves en bruger, så man undgår at være root. Kommandoerne kommer her (antager at brugeren skal hedde `manager`)

```bash
$ adduser manager
$ passwd manager
$ usermod -aG wheel manager
$ rsync --archive --chown=manager:manager ~/.ssh /home/manager
```

Man kan skifte til den nyoprettede bruger med `su - manager` eller forlade SSH-sessionen og bruge `manager@<ip-address>`.

# Pakker
Dette projekt kræver at følgende pakker installeres: `git nginx nano`. Det kan gøres med

```bash
$ sudo dnf -y install git nginx nano
```

## Nginx
Lad os begynde med at starte nginx serveren. Man kunne også have brugt [Apache HTTPD](https://httpd.apache.org/), men jeg har lidt erfaring med nginx, da jeg har brugt det før. Meget lidt er klart nu, men det vil være rart at se noget i browseren og få en fornemmelse for, at der sker noget. Så lad os starte for den:

```bash
$ sudo systemctl enable --now nginx
```

Man vil nu være i stand til at se ens hjemmeside ved at gå til `http://<ip-address>/`. Fantastisk! Dog mangler hængelåsen. Lad os også fikse det, nu mens vi er her. 

### Certbot

Jeg har brugt Lets Encrypt og deres certbot til at generere certifikaterne. certbot kan installeres med

```bash
$ sudo dnf install epel-release
$ sudo dnf install certbot python3-certbot-nginx
```

Genererer certikaterne ved at køre (mere information på [certbot.eff.org](https://certbot.eff.org/instructions?ws=nginx&os=centosrhel8))

```bash
$ sudo certbot --nginx
$ sudo certbot renew --dry-run
```

## Hugo
Bloggen skrives i .md filer, hvor [Hugo](https://gohugo.io/) bygger hjemmesiden hver gang man pusher en ny fil til ens GitHub repository (i dette tilfælde). Andre programmer man kunne bruge, der løser samme problem som Hugo, er bl.a. [Eleventy](https://www.11ty.dev/), [Zola](https://www.getzola.org/) og [Jekyll](https://jekyllrb.com/). Zola ser også lovende ud, men det endte med Hugo i denne omgang.

Det betyder, at man nu er klar til at gå videre og installere Hugo. Det nemmeste at gøre, hvis man bruger Rocky Linux er at installere snap, og så installere Hugo derfra. (Man behøver ikke køre den første linje, hvis man gjorde det før).

```bash
$ sudo dnf install epel-release
$ sudo dnf install snapd
$ sudo systemctl enable --now snapd.socket
$ sudo ln -s /var/lib/snapd/snap /snap
```

Jeg følger guiden fra [snapcraft.io](https://snapcraft.io/docs/installing-snap-on-rocky), og den anbefaler, at man genstarter serveren eller logger ud?. Man kan genstarte serveren med `sudo reboot` og endelig installere Hugo med:

```bash
$ snap install hugo
```

Fedt. Lav en hjemmeside med

```bash
$ hugo new site <sidens navn>
```

Man kan finde temaer fra [themes.gohugo.io](https://themes.gohugo.io/). Denne blog bruger `papermod`.

```bash
$ cd <sidens navn>
$ git init
$ git sobmodule add https://github.com/adityatelange/hugo-PaperMod.git themes/papermod
$ echo theme = \"ananke\" >> config.toml
```

I `config.toml` kan bl.a. sprog, sidens titel, URL osv ændres. Det første indlæg laves med

```bash
$ hugo new posts/my-first-post.md
```

Ændr filen og giv den indhold. Hjemmesiden bygges med

```bash
$ hugo
```

Hvis den selv skal opdatere, når indholdet ændres, kan man bruge

```bash
$ hugo server
```

Tilføj `-D` til kommandoerne for at indlæg markeret som kladder også inkluderes.

# Automatisk bygning af hjemmeside
Bloggen kan findes på [GitHub](https://github.com/NikolajK-HTX/njk-blog). Hver gang man pusher et nyt commit til mit GitHub repository, er en Webhook sat op til at sende en POST request til serveren på DO.

Der har jeg en Go-webserver, der lytter bag nginx via reverse proxy, som kører et shell script, som bygger og opdaterer siden.

```bash
$ #!/bin/bash
$ cd "$(dirname "$0")"
$ git pull
$ hugo
$ rsync -avu --delete public/ /usr/share/nginx/blog
```

I stedet for at have en dedikeret webserver til at køre et shell script, burde man bruge [CGI](https://en.wikipedia.org/wiki/Common_Gateway_Interface). Problemet er bare, at det understøtter "standard" nginx ikke (men det gør Apache HTTPD). Nginx har noget andet man kan bruge: [FastCGI](https://www.nginx.com/resources/wiki/start/topics/examples/fastcgiexample/).

(Fordi Nginx bliver brugt som en reverse proxy til en anden webserver på samme maskine, får man lidt bøvl med SELinux. Som standard har nginx ikke adgang til andre porte end 80 og 443, det ændres med `setsebool -P httpd_can_network_connect 1`.
