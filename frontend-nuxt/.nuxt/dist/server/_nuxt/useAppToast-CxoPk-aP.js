import { E as useToast } from "../server.mjs";
const useAppToast = () => {
  const toast = useToast();
  const success = (message) => {
    toast.add({ severity: "success", summary: "Success", detail: message, life: 3e3 });
  };
  const error = (message) => {
    toast.add({ severity: "error", summary: "Error", detail: message, life: 5e3 });
  };
  const info = (message) => {
    toast.add({ severity: "info", summary: "Info", detail: message, life: 3e3 });
  };
  const warning = (message) => {
    toast.add({ severity: "warn", summary: "Warning", detail: message, life: 4e3 });
  };
  return {
    success,
    error,
    info,
    warning
  };
};
export {
  useAppToast as u
};
//# sourceMappingURL=useAppToast-CxoPk-aP.js.map
