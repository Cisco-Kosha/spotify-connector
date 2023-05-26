# Kosha Spotify Connector

Spotify is a popular music streaming platform that provides APIs and developer tools that enable you to integrate with its platform to build music-related applications and create personalized experiences. 

The Kosha Spotify connector enables you to perform REST API operations from the Spotify API in your Kosha workflow or custom application. Using the Kosha Spotify connector, you can directly access the Spotify platform to:

* Get a user's profile, playlist, and followers
* Get a user's saved albums, tracks, and audiobooks
* Create and add items to playlists
* Control the Spotify player

## Useful Actions

You can use the Kosha Spotify connector to manage Spotify albums, tracks, users, playlists, libraries, and more.

Refer to the Kosha Spotify connector [API specification](openapi.json) for details.

### Albums

Use the albums API to:

* Get a specific album or several albums
* Get a user's saved albums or save albums for a user
* Get new releases

### Tracks 

Use the tracks API to:

* Get a specific track or several tracks
* Get a user's saved tracks or save tracks for a user
* Get recommended tracks

### Users

Use the users API to:

* Get a user's profile, followed artists, and playlists
* Check if a user follows a particular artist or playlist

### Playlists

Use the playlists API to:

* Create and add items to a playlist
* Follow or unfollow playlists
* Get feature playlists

## Authentication

Kosha uses OAuth 2.0 to connect to Spotify. If you already have an application registered with Spotify, you can use your `client_id` and `client_secret` when provisioning the connector.

If you don’t want to use those credentials when provisioning the Spotify connector, Kosha provides bootstrap credentials. After you sign in to your Spotify app, Spotify gives Kosha an access token and your connector is provisioned. Kosha automatically refreshes your access token before it expires to ensure there’s no disruption in use.

## Kosha Connector Open Source Development

All connectors Kosha shares on the marketplace are open source. We believe in fostering collaboration and open development. Everyone is welcome to contribute their ideas, improvements, and feedback for any Kosha connector. We encourage community engagement and appreciate any contributions that align with our goals of an open and collaborative API management platform.

Refer to the contribution guidelines for details.