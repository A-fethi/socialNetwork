import router from "../router";

export const handleAuthError = (error) => {
  if (
    error.response &&
    (error.response.status === 401 || error.response.status === 403)
  ) {
    console.error("Authentication error:", error);
    router.push("/login");
    return true;
  }
  return false;
};

export const authenticatedFetch = async (url, options = {}) => {
  try {
    const fetchOptions = {
      ...options,
      credentials: "include",
    };

    const response = await fetch(url, fetchOptions);

    if (response.status === 401 || response.status === 403) {
      console.error("Authentication error:", response.statusText);
      router.push("/login");
      throw new Error("Authentication failed");
    }

    return response;
  } catch (error) {
    handleAuthError(error);
    throw error;
  }
};
