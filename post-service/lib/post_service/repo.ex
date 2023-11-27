defmodule PostService.Repo do
  use Ecto.Repo,
    otp_app: :post_service,
    adapter: Ecto.Adapters.Postgres
end
