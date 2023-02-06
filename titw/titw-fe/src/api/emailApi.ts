import axios from "axios";

const emailApi = axios.create({
  baseURL: "http://localhost:3000",
});

export default emailApi;
