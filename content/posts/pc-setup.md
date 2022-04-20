---
title: "PC Setup"
date: 2022-04-20T12:24:00Z
ShowToc: true
draft: false
---

# Ubuntu

Disse pakker kommer fra Ubuntus depoter. 

```bash
$ sudo apt update
$ sudo apt install git speedcrunch ghex texlive texstudio lz4 postgresql rabbitmq-server golang build-essential htop zeal curl
```

## Snap

Disse pakker kommer fra snapstore. Man behøver ikke bruge snap man kan også bruge flatpak. Nogle eksempler er under Fedora pakker.

```bash
$ snap install code --classic
$ snap install sublime sublime-text --classic
$ snap install discord spotify
$ snap install blender --classic
$ snap install krita inkscape kdenlive gimp
```

# Fedora

Kommandoerne er taget fra [https://rpmfusion.org/Configuration](rpmfusion.org). Rul ned til overskriften *Command Line Setup using rpm*.

```bash
$ sudo dnf install https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
```

Kommandoerne er taget fra følgende side: [https://docs.fedoraproject.org/en-US/quick-docs/assembly_installing-plugins-for-playing-movies-and-music/](docs.fedoraproject.org)

```bash
$ sudo dnf install gstreamer1-plugins-{bad-\*,good-\*,base} gstreamer1-plugin-openh264 gstreamer1-libav --exclude=gstreamer1-plugins-bad-free-devel
$ sudo dnf install lame\* --exclude=lame-devel
$ sudo dnf group upgrade --with-optional Multimedia
```

```bash
$ sudo dnf -y install htop foliate git apostrophe ghex zeal epiphany texstudio texlive-scheme-medium texlive-babel-danish texlive-hyphen-danish remmina golang curl rabbitmq-server postgresql-server
$ sudo dnf -y groupinstall "Development Tools"
```

## Flatpak

Der kan forekomme problemer med pakkerne: Atom, Blender og Code. Mange af dem er ikke officielle, hvorfor de ikke virker 100%. De fleste er officielt udgivet med snap. Hvis man ønsker at bruge snap i stedet for kan man springe nedenstående segment over - installere snap og følge instrukserne under Ubuntu pakker.

```bash
$ flatpak install flathub io.atom.Atom
$ flatpak install flathub com.spotify.Client
$ flatpak install flathub com.discordapp.Discord
$ flatpak install flathub org.blender.Blender
$ flatpak install flathub org.kde.kdenlive
$ flatpak install flathub org.kde.krita
$ flatpak install flathub com.visualstudio.code
```
Instrukser for at installere snap på Fedora. De er taget fra [https://snapcraft.io/docs/installing-snap-on-fedora](snapcraft.io).

```bash
$ sudo dnf install snapd # Log ud eller genstart efter denne kommando
$ sudo ln -s /var/lib/snapd/snap /snap # Log ud eller genstart igen
```

# Andre

## Setup rust-lang

```bash
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## Aspell

[http://aspell.net/](aspell.net) er en samling af ordbøger. ``aspell-da`` er den danske ordbog.
