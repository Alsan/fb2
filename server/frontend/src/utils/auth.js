import store from "@/store";
import router from "@/router";
import { Base64 } from "js-base64";
import { baseURL } from "@/utils/constants";
import { md5, bcrypt } from "hash-wasm";

export async function md5Hash(password) {
  return await md5(password.split("").reverse().join());
}

export async function bcryptHash(password) {
  const salt = new Uint8Array(16);
  crypto.getRandomValues(salt);

  return bcrypt({ password, salt, costFactor: 11, outputType: "encoded" });
}

export function parseToken(token) {
  const parts = token.split(".");

  if (parts.length !== 3) {
    throw new Error("token malformed");
  }

  const b64Parts = Base64.decode(parts[1]);
  console.log(`b64Parts: ${b64Parts}`);
  const data = JSON.parse(b64Parts);

  document.cookie = `auth=${token}; path=/`;

  localStorage.setItem("jwt", token);
  store.commit("setJWT", token);
  store.commit("setUser", data.user);
}

export async function validateLogin() {
  try {
    if (localStorage.getItem("jwt")) {
      await renew(localStorage.getItem("jwt"));
    }
  } catch (_) {
    console.warn("Invalid JWT token in storage"); // eslint-disable-line
  }
}

export async function login(username, password, recaptcha) {
  const data = { username, password, recaptcha };

  const res = await fetch(`${baseURL}/api/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });

  const body = await res.text();

  if (res.status === 200) {
    parseToken(body);
  } else {
    throw new Error(body);
  }
}

export async function renew(jwt) {
  const res = await fetch(`${baseURL}/api/renew`, {
    method: "POST",
    headers: {
      "X-Auth": jwt
    }
  });

  const body = await res.text();

  if (res.status === 200) {
    parseToken(body);
  } else {
    throw new Error(body);
  }
}

export async function signup(username, password) {
  const data = { username, password };

  const res = await fetch(`${baseURL}/api/signup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });

  if (res.status !== 200) {
    throw new Error(res.status);
  }
}

export function logout() {
  document.cookie = "auth=; expires=Thu, 01 Jan 1970 00:00:01 GMT; path=/";

  store.commit("setJWT", "");
  store.commit("setUser", null);
  localStorage.setItem("jwt", null);
  router.push({ path: "/login" });
}
