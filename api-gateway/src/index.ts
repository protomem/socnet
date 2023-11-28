import { Elysia } from "elysia";

const app = new Elysia().get("/", () => "Hello Elysia").listen(8081);

console.log(
  `API-Gateway is running at ${app.server?.hostname}:${app.server?.port}`,
);
