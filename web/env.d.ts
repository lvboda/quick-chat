/// <reference types="vite/client" />

interface ImportMetaEnv {
  VITE_MODE_NAME: string;
  VITE_LOGIN_TEST: string;
  VITE_HTTP_URL: string;
  VITE_WS_URL: string;
  VITE_APP_TITLE: string;
}
interface Window {
  SimplePeer: any;
}
