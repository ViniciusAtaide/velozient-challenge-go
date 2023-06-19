import { useCallback, useMemo, useReducer, useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { useMutation, useQueryClient } from "react-query";

export type IPasswordCard = {
  id: string;
  url: string;
  name: string;
  username: string;
  password: string;
};

type Props = {
  card: IPasswordCard;
};

type State = Partial<Record<keyof IPasswordCard, boolean>>;

type ActionTypes = "ToggleEditOn";

type Action = {
  type: ActionTypes;
  payload: keyof IPasswordCard;
};

const reducer = (state: State, action: Action) => {
  switch (action.type) {
    case "ToggleEditOn": {
      return {
        ...state,
        [action.payload]: !state[action.payload],
      };
    }
  }
};

export default function PasswordCard({ card }: Props) {
  // I could move all this to a custom hook, depends on the size, this is still bearable to me
  const [hidden, setHidden] = useState(false);
  const {
    register,
    getValues,
    handleSubmit,
    reset,
    formState: { isDirty },
  } = useForm({ defaultValues: card });
  const password = getValues("password");

  const hiddenPassword = useMemo(
    () => (hidden ? "*********" : password),
    [hidden, password]
  );

  const [editable, dispatch] = useReducer(reducer, {});

  const client = useQueryClient();

  const updateCard = useMutation<unknown, unknown, IPasswordCard>({
    mutationFn: (card) =>
      fetch(`${import.meta.env.VITE_BACKEND_URL}/password-cards/${card.id}`, {
        method: "PATCH",
        body: JSON.stringify(card),
        headers: {
          "Content-Type": "application/json",
        },
      }).then((resp) => {
        if (resp.status >= 400) {
          throw Error();
        }
      }),
    onError: () => {
      reset();
    },
  });

  const toggleEditField = useCallback(
    (field: keyof IPasswordCard) => {
      dispatch({ type: "ToggleEditOn", payload: field });
    },
    [dispatch]
  );

  const onUpdate: SubmitHandler<IPasswordCard> = useCallback(
    (card) => {
      if (isDirty) {
        updateCard.mutate(card);
        reset(card, { keepValues: true });
      }

      Object.keys(editable)
        .filter((e) => !!editable[e as keyof IPasswordCard])
        .forEach((e) => toggleEditField(e as keyof IPasswordCard));
    },
    [updateCard, isDirty, toggleEditField, editable, reset]
  );

  const deleteCard = useMutation<unknown, unknown, string>({
    mutationFn: (id) =>
      fetch(`${import.meta.env.VITE_BACKEND_URL}/password-cards/${id}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
      }),

    onSuccess: () => {
      // I Could've updated only the client cache removing the password-card without needing to do a refetch, this is a simplification
      client.invalidateQueries("password-cards");
    },
  });

  return (
    <ul className="flex p-5 flex-col border-2 space-y-3  hover:bg-blue-100 transition-all">
      <div className="my-3 flex space-x-5 items-center">
        <label
          htmlFor="hidden"
          onClick={() => setHidden((h) => !h)}
          className="flex-1 gap-2 text-left"
        >
          <input
            className="text-lg"
            name="hidden"
            checked={hidden}
            type="checkbox"
          />
          Hide Password?
        </label>
        {editable.name ? (
          <input
            onMouseLeave={handleSubmit(onUpdate)}
            type="text"
            className="h-8 w-32 border-2 box-border border-gray-800 text-end"
            id="name"
            placeholder="Name"
            {...register("name")}
          />
        ) : (
          <h1
            onMouseEnter={() => toggleEditField("name")}
            className="text-right h-8 w-32 font-medium text-lg"
          >
            {getValues("name")}
          </h1>
        )}
        <button onClick={() => deleteCard.mutate(getValues("id"))}>X</button>
      </div>
      <li className="flex place-content-between h-7">
        <label>Username:</label>
        {editable.username ? (
          <span>
            <input
              onMouseLeave={handleSubmit(onUpdate)}
              {...register("username")}
              className="b-1 border-2 box-border border-gray-800 text-end px-2"
              type="text"
              name="username"
              id="username"
              placeholder="Username"
            />
          </span>
        ) : (
          <span onMouseEnter={() => toggleEditField("username")}>
            {getValues("username")}
          </span>
        )}
      </li>
      <li className="flex place-content-between h-7">
        <label>Password: </label>
        {editable.password ? (
          <span>
            <input
              onMouseLeave={handleSubmit(onUpdate)}
              {...register("password")}
              className="b-1 border-2 box-border border-gray-800 text-end px-2"
              name="password"
              id="password"
              type={hidden ? "password" : "text"}
              placeholder="Password"
            />
          </span>
        ) : (
          <span onMouseEnter={() => toggleEditField("password")}>
            {hiddenPassword}
          </span>
        )}
      </li>
      <li className="flex place-content-between h-7">
        <label>Url:</label>
        {editable.url ? (
          <span>
            <input
              onMouseLeave={handleSubmit(onUpdate)}
              {...register("url")}
              className="p-1 box-border border-2 border-gray-800 text-end px-2"
              type="url"
              name="url"
              id="url"
              placeholder="URL"
            />
          </span>
        ) : (
          <span onMouseEnter={() => toggleEditField("url")}>
            {getValues("url")}
          </span>
        )}
      </li>
    </ul>
  );
}
