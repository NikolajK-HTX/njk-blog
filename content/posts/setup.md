---
title: "Setup blog"
date: 2022-04-10T18:39:35Z
draft: false
---

Denne post er til for at hjælpe med at sætte sådant et projekt op.

Jeg går ud fra, at man bruger Rocky Linux eller tilsvarende (det kunne f.eks. AlmaLinux, CentOS eller Fedora Server). Det vil dog ikke se voldsomt anderledes ud, hvis man valgte at sætte det op på en anden distro.

Du læser højst sandsynligt dette fra [blog.jehaj.dk](https://blog.jehaj.dk/), som bliver leveret til dig fra en server hos DigitalOcean.

# Sæt igang
Du kan ikke gå igang før, du har fået adgang til din server. Det gør du med `ssh root@<ip-address>`. I stedet for `<ip-address>` kan du bruge dit domæne, hvis du har sat det til at pege på serveren. DigitalOcean anbefaler man bruger SSH passphrases, så det gør vi.

(Det er ikke sikkert, at maskinen er helt opdateret... `dnf upgrade`)

Dernæst laves en bruger, så vi undgår at være root. Kommandoerne kommer her (antager at brugeren skal hedde `manager`):
```bash
$ adduser manager
$ passwd manager
$ usermod -aG wheel manager
$ rsync --archive --chown=manager:manager ~/.ssh /home/manager
```
Du kan skifte til den nyoprettede bruger med `su - manager` eller forlade SSH-sessionen og bruge `manager@<ip-address>`.

# Pakker
Dette projekt kræver at følgende pakker installeres: `git nginx nano`. Det kan gøres med
```bash
$ sudo dnf -y install git nginx nano
```

## Nginx
Lad os starte med at starte nginx serveren. Meget lidt er klart nu, men det vil være rart at se noget i browseren og få en fornemmelse for, at der sker noget. Så lad os starte for den:

```bash
$ sudo systemctl enable --now nginx
```

Du vil nu være i stand til at se din hjemmeside ved at gå til `http://<ip-address>/`. Fantastisk! Dog mangler hængelåsen. Lad os også fikse det, nu mens vi er her. Jeg har brugt Lets Encrypt og deres certbot til at generere certifikaterne. `certbot` kan installeres med

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
Bloggen skrives i .md filer, hvor [Hugo](https://gohugo.io/) bygger hjemmesiden hver gang man pusher en ny fil til ens GitHub repository (i dette tilfælde).

Det betyder, at vi nu skal installere Hugo. Det nemmeste at gøre, hvis man bruger Rocky Linux er at installere snap, og så installere Hugo derfra. (Du behøver ikke køre den første linje, hvis du gjorde det før).

```bash
$ sudo dnf install epel-release
$ sudo dnf install snapd
$ sudo systemctl enable --now snapd.socket
$ sudo ln -s /var/lib/snapd/snap /snap
```

Jeg følger guiden fra [snapcraft.io](https://snapcraft.io/docs/installing-snap-on-rocky), og den anbefaler, at du genstarter serveren eller logger ud?. Du kan genstarte serveren med `sudo reboot` og endelig installere Hugo med:

```bash
$ snap install hugo
```
