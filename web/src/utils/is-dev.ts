function isDev(): boolean {
  return import.meta.env.VITE_MODE_NAME === "development";
}

export default isDev;
