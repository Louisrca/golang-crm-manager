# CRM Manager

Bienvenue sur CRM Manager, une application en ligne de commande pour la gestion simplifiée de vos contacts utilisateurs.

> **Note** : Chaque release est identifiée par un tag `v0.x` où `x` correspond au jour du cours.

## Fonctionnalités

- Ajouter un utilisateur
- Lister tous les utilisateurs
- Modifier un utilisateur existant
- Supprimer un utilisateur
- Ajouter un contact rapidement via flags CLI

## Prérequis

- Go 1.x ou supérieur

## Installation

```bash
git clone <url-du-repo>
cd crm-manager
```

## Utilisation

### Lancer l'application en mode interactif

```bash
go run cmd/api/main.go
```

L'application affichera un menu interactif pour naviguer entre les différentes fonctionnalités.

### Ajouter un utilisateur via la CLI

Vous pouvez ajouter un utilisateur directement en ligne de commande avec des flags :

```bash
go run cmd/api/main.go -name jeanmich -email jm@mail.fr
```

### Aide

Pour découvrir tous les flags disponibles et leurs options :

```bash
go run cmd/api/main.go --help
```

## Exemples d'utilisation

```bash
# Ajouter un contact
go run cmd/api/main.go -name jeanmich -email jeanmich@mail.com

# Lancer le menu interactif
go run cmd/api/main.go
```

## Structure du projet

```
crm-manager/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── domain/
│   utils/
└── README.md
```
