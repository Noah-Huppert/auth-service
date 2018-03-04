# Auth Service
Simple authentication service which provides proof of identity to other services.

# Table Of Contents
- [Overview](#overview)

# Overview
Auth service is a common authentication service which uses the 
[Vault Project](https://www.vaultproject.io) as a authentication backend.  

Auth service provides the following features:  

- Login user interface
	- [User & Password](https://www.vaultproject.io/docs/auth/userpass.html)
	- [(WIP) Time Based One Time Passwords](https://www.vaultproject.io/docs/secrets/totp/index.html)
	- [(WIP) LDAP](https://www.vaultproject.io/docs/auth/ldap.html)
	- [(WIP) GitHub](https://www.vaultproject.io/docs/auth/github.html)
- [Vault Secret Engine API](https://www.vaultproject.io/docs/secrets/index.html)
	- Provision over 12 different types of secrets for users which 
	  authenticate with auth service
- LDAP server
	- Provide federated access to users
- Roll based access control
- User management dashboard
- [Vault Audit log](https://www.vaultproject.io/docs/audit/index.html) user interface
