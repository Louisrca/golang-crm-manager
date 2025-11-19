# CRM Manager

Bienvenue sur CRM Manager, une application en ligne de commande pour la gestion simplifiée de vos contacts utilisateurs avec support de multiples systèmes de stockage.

> **Note** : Chaque release est identifiée par un tag `v0.x` où `x` correspond au jour du cours.

## Fonctionnalités

- Ajouter un utilisateur
- Lister tous les utilisateurs
- Modifier un utilisateur existant
- Supprimer un utilisateur
- Support de 3 types de stockage : JSON, GORM (SQLite), ou Mémoire locale
- Configuration flexible via fichier YAML ou flags CLI

## Prérequis

- Go 1.x ou supérieur

## Installation

```bash
git clone <url-du-repo>
cd crm-manager
go build .
```

## Configuration

### Fichier config.yaml

```yaml
store: JSON
json_path: users.json
db_path: users.db
```

Si aucun fichier de configuration n'est présent, l'application utilisera les valeurs par défaut.

## Utilisation

### Commandes disponibles

#### Ajouter un utilisateur

```bash
./golang-crm-manager add --store JSON
```

#### Lister les utilisateurs

```bash
./golang-crm-manager users-list --store JSON
```

#### Mettre à jour un utilisateur

```bash
./golang-crm-manager update-user --store JSON
```

#### Supprimer un utilisateur

```bash
./golang-crm-manager delete-user --store JSON
```

### Options de stockage

Vous pouvez spécifier le type de stockage via le flag `--store` :

- `JSON` : Stockage dans un fichier JSON (par défaut)
- `GORM` : Base de données SQLite via GORM
- `LOCAL` : Stockage en mémoire (temporaire)

### Exemples d'utilisation

```bash
# Utiliser le stockage JSON
./golang-crm-manager add --store JSON

# Utiliser le stockage GORM (SQLite)
./golang-crm-manager users-list --store GORM

# Utiliser le stockage en mémoire
./golang-crm-manager add --store LOCAL

# Lister les utilisateurs avec la config par défaut
./golang-crm-manager users-list
```

### Aide

Pour découvrir tous les flags disponibles et leurs options :

```bash
./golang-crm-manager --help
```

## Architecture

Le projet utilise :

- **Cobra** pour la gestion des commandes CLI
- **Viper** pour la configuration (YAML + flags)
- **GORM** pour l'accès aux bases de données
- **Interface Storer** pour une abstraction du stockage

## Dépendances

```bash
go get github.com/spf13/cobra
go get github.com/spf13/viper
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```
