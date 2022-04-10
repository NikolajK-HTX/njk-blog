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
```
$ adduser manager
$ passwd manager
$ usermod -aG wheel manager
$ rsync --archive --chown=manager:manager ~/.ssh /home/manager
```
Du kan skifte til den nyoprettede bruger med `su - manager` eller forlade SSH-sessionen og bruge `manager@<ip-address>`.

# Pakker
Dette projekt kræver at følgende pakker installeres: `git nginx nano`. Det kan gøres med
```
$ sudo dnf -y install git nginx nano
```

Bloggen skrives i .md filer, hvor [Hugo](https://gohugo.io/) bygger hjemmesiden hver gang man pusher en ny fil til ens GitHub repository (i dette tilfælde).

Det betyder, at vi nu skal installere Hugo. Det nemmeste at gøre, hvis man bruger Rocky Linux er at installere snap, og så installere Hugo derfra.

```
$ sudo dnf install epel-release
$ sudo dnf install snapd
$ sudo systemctl enable --now snapd.socket
$ sudo ln -s /var/lib/snapd/snap /snap
```

Jeg føler guiden fra [snapcraft.io](https://snapcraft.io/docs/installing-snap-on-rocky) og den anbefaler, at du genstarter serveren eller logger ud?. Du kan genstarte serveren med `sudo reboot` og endelig installere Hugo med:

```
$ snap install hugo
```
