import { decodeBase64Url } from "./helpers.js";

export async function createCredentials() {
  const credentails = await navigator.credentials.create({
    publicKey: {
      rp: {
        name: "a10a server",
      },
      user: {
        id: decodeBase64Url("some-user-id"),
        name: "mail@a10a.app",
        displayName: "mail@a10a.app",
      },
      challenge: decodeBase64Url("some-challenge"),
      pubKeyCredParams: [
        {
          type: "public-key",
          alg: -7,
        },
        {
          type: "public-key",
          alg: -257,
        },
      ],
    },
  });

  console.log({ credentails });
}
