import { useMutation, useQueryClient } from "react-query";
import { useForm, SubmitHandler, RegisterOptions } from "react-hook-form";
import { useCallback } from "react";

type FormInputs = {
  username: string;
  name: string;
  password: string;
  url: string;
};

const urlPattern = new RegExp(
  /^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_+.~#?&/=]*)$/
);

const messages = (field: string) => ({
  required: `${field} is required`,
  min: (min: number) => `${field} must be at least ${min} chars long`,
  max: (max: number) => `${field} must be no longer than ${max}`,
  pattern: `${field} must be a valid url`,
});

const validation = {
  username: {
    required: {
      value: true,
      message: messages("Username").required,
    },
    minLength: {
      value: 3,
      message: messages("Username").min(3),
    },
    maxLength: {
      value: 32,
      message: messages("Username").max(32),
    },
  } as RegisterOptions<FormInputs, "username">,
  name: {
    required: {
      value: true,
      message: messages("Name").required,
    },
    minLength: {
      value: 3,
      message: messages("Name").min(3),
    },
  } as RegisterOptions<FormInputs, "name">,
  password: {
    required: { value: true, message: messages("Password").required },
    minLength: { value: 8, message: messages("Password").min(8) },
  } as RegisterOptions<FormInputs, "password">,
  url: {
    required: { value: true, message: messages("Url").required },
    pattern: { value: urlPattern, message: messages("Url").pattern },
  } as RegisterOptions<FormInputs, "url">,
} as const;

export default function CreatePasswordModal() {
  const client = useQueryClient();
  const createCard = useMutation<unknown, unknown, FormInputs>({
    mutationFn: (card) =>
      fetch(`${import.meta.env.VITE_BACKEND_URL}/password-cards`, {
        method: "POST",
        body: JSON.stringify(card),
        headers: {
          "Content-Type": "application/json",
        },
      }),
    onSuccess: () => {
      client.invalidateQueries("password-cards");
      dismiss();
    },
    onError: (error) => {
      console.log(error);
    },
  });
  const dismiss = useCallback(
    () => client.setQueryData("modal", () => false),
    [client]
  );
  const onSubmit: SubmitHandler<FormInputs> = useCallback(
    async (data) => {
      createCard.mutate(data);
    },
    [createCard]
  );

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormInputs>();

  return (
    <div
      className="relative z-10"
      aria-labelledby="modal-title"
      role="dialog"
      aria-modal="true"
    >
      <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>

      <div className="fixed inset-0 z-10 overflow-y-auto">
        <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-[500px]">
              <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                <div className="sm:flex sm:items-start">
                  <div className="mt-3 flex-1 text-center sm:mt-0 sm:ml-4 sm:text-left">
                    <h3
                      className="text-lg mb-10 text-center font-medium leading-6 text-gray-900"
                      id="modal-title"
                    >
                      Create new Password Card
                    </h3>
                    <div className="mt-2 flex flex-col gap-3">
                      <div className="flex-1 flex align-items-bottom">
                        <label htmlFor="name" className="flex-1 text-lg">
                          Name:
                        </label>
                        <div className="flex-1">
                          <input
                            {...register("name", validation["name"])}
                            type="text"
                            placeholder="Google"
                            className="box-border text-center border-2 p-1 border-gray-700 rounded-md"
                            name="name"
                            id="name"
                          />
                          {errors.name && (
                            <div className="text-red-600">
                              <div>{errors.name.message}</div>
                            </div>
                          )}
                        </div>
                      </div>
                      <div className="flex-1 flex align-items-bottom">
                        <label htmlFor="username" className="flex-1 text-lg">
                          Username:
                        </label>
                        <div className="flex-1">
                          <input
                            {...register("username", validation["username"])}
                            placeholder="Username"
                            type="text"
                            className="box-border text-center border-2 p-1 border-gray-700 rounded-md"
                            name="username"
                            id="username"
                          />
                          {errors.username && (
                            <div className="text-red-600">
                              <div>{errors.username.message}</div>
                            </div>
                          )}
                        </div>
                      </div>
                      <div className="flex-1 flex align-items-bottom">
                        <label htmlFor="password" className="flex-1 text-lg">
                          Password:{" "}
                        </label>
                        <div className="flex-1">
                          <input
                            {...register("password", validation["password"])}
                            placeholder="Password"
                            type="password"
                            className="box-border text-center border-2 p-1 border-gray-700 rounded-md"
                            name="password"
                            id="password"
                          />
                          {errors.password && (
                            <div className="text-red-600">
                              <div>{errors.password.message}</div>
                            </div>
                          )}
                        </div>
                      </div>
                      <div className="flex-1 flex align-items-bottom">
                        <label htmlFor="url" className="flex-1 text-lg">
                          URL:
                        </label>
                        <div className="flex-1">
                          <input
                            {...register("url", validation["url"])}
                            placeholder="https://www.google.com"
                            type="text"
                            className="box-border text-center border-2 p-1 border-gray-700 rounded-md"
                            name="url"
                            id="url"
                          />
                          {errors.url && (
                            <div className="text-red-600">
                              <div>{errors.url.message}</div>
                            </div>
                          )}
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div className="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                <button
                  type="submit"
                  className="inline-flex disabled:bg-gray-400 w-full justify-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm"
                  disabled={createCard.isLoading}
                >
                  {createCard.isLoading ? "Submitting" : "Submit"}
                </button>
                <button
                  type="button"
                  className="mt-3 inline-flex w-full justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-base font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
                  onClick={dismiss}
                >
                  Cancel
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
