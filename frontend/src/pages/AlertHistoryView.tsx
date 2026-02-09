import { useEffect, useState } from "react";
import { useNavigate, useParams, } from "react-router";
import type { AlertHistory } from "../types/Alert";
import { getHistoryAlert } from "../api/GetHistoryAlert";
import AlertHistoryTable from "../components/AlertHistoryTable";
import { formatAlertInstance } from "../utils/Formatter";
import './AlertHistoryView.css';

export default function AlertHistoryView() {

  const { alertInstance } = useParams();
  const [history, setHistory] = useState<AlertHistory>();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    async function loadAlertHistory() {

      let showLoading: any;

      try {
        showLoading = setTimeout(() => {
          setLoading(true);

        }, 500);
        const data = await getHistoryAlert(alertInstance!);
        setHistory(data);

      } catch (error: any) {
        setError(error.message)
      }
      finally {
        clearTimeout(showLoading);
        setLoading(false);
      }
    }
    if (!history) {
      loadAlertHistory()
    }
  }, [alertInstance]);

  if (loading) return <h1>Lädt History von {alertInstance}</h1>;
  if (error) return <h1>Fehler: {error}</h1>;
  if (!history) return;
  if (!history.events) return <h1>Keine History für {alertInstance} gefunden</h1>

  return (
    <>
      <div onClick={() => navigate("/")} style={{ display: "flex" }}>
        <button style={{ marginRight: "auto" }}>BACK</button>
      </div>
      <h1>{formatAlertInstance(history.instance)}</h1>
      < AlertHistoryTable alertHistory={history} ></AlertHistoryTable >
    </>
  )

}
