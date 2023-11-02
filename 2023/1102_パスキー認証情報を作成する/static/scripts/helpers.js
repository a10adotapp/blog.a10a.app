export function encodeBase64Url(buffer) {
  return btoa(String.fromCharCode(...new Uint8Array(buffer)))
    .replace(/=/g, "")
    .replace(/\+/g, "-")
    .replace(/\//g, "_");
}

export function decodeBase64Url(str) {
  return Uint8Array.from(str, (char) => (char.charCodeAt(0)));
}
