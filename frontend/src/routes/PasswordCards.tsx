import CreatePasswordModal from "@/components/CreatePasswordModal";
import PasswordCard, { IPasswordCard } from "@/components/PasswordCard";
import { useState } from "react";
import { useQuery, useQueryClient } from "react-query";
import { useDebounce } from "usehooks-ts";

export default function PasswordCards() {
  // global variable, we could use zustand for this purpose but this project is too small
  const { data: modal } = useQuery("modal", () => false);
  const [search, setSearch] = useState("");
  const client = useQueryClient();

  const debouncedSearch = useDebounce(search, 1000);

  const response = useQuery<IPasswordCard[], string>(
    ["password-cards", debouncedSearch],
    () =>
      fetch(
        `${
          import.meta.env.VITE_BACKEND_URL
        }/password-cards?limit=0&name=${debouncedSearch}`
      ).then((resp) => resp.json()),
    {}
  );

  if (response.error) return <>Error connecting to server...</>;

  return (
    <>
      <div className="flex flex-col gap-5 w-96">
        {modal && <CreatePasswordModal />}

        <div className="flex flex-1 gap-2">
          <input
            name="search"
            className="border-2 flex-1 p-2 box-border text-center"
            placeholder="Search for a card"
            type="text"
            onChange={(event) => setSearch(event.target.value)}
            id="search"
            autoFocus
          />

          <button
            onClick={() => client.setQueryData("modal", () => true)}
            className="text-lg w-10  bg-blue-700 text-white rounded-sm"
          >
            +
          </button>
        </div>

        {response.isLoading || !response.data ? (
          <>Loading...</>
        ) : (
          <>
            <div className="container flex flex-col gap-4 box-border">
              {response.data.map((card) => (
                <PasswordCard key={card.id} card={card} />
              ))}
              {response.data.length === 0 && <>No cards to show</>}
            </div>
          </>
        )}
      </div>
    </>
  );
}
