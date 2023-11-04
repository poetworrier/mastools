# mastools

A bunch of tools

## Trends

See <https://docs.joinmastodon.org/methods/admin/trends/>
There are three trend types:

- Links
- Statuses
- Tags

To see how it works look at the out from `git grep "manage_taxonomies"` and `git grep "Admin::Trends"`.

NOTE: I personally disable links on servers I admin. There is too much risk from a bad/malicious link.

### Handling Permissions

It's never a good idea to have a lot of credentials floating around with broad permissions. To set
 this up soundly:

- Create a "Trend Editor" role that can ONLY "Manage Taxonomies"
- Create a service account and grant it that role
- Login to this account, setup 2FA, etc.
- Create a new application on the service account
  - Revoke all permissions except "admin:write". Unfortunately Mastodon doesn't have a fine grained permission for "admin:trends".
- Copy the access token into a secret manager.
- Pass the token to this application when calling methods
