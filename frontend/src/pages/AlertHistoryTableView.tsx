import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router';
import type { AlertHistory } from '../types/Alert';
import { getAlertHistory } from '../api/GetAlertHistory';
import AlertHistoryTable from '../components/AlertHistoryTable';
import { formatAlertInstance } from '../utils/Formatter';

export default function AlertHistoryView() {
    const { alertInstance } = useParams();
    const [alertHistory, setAlertHistory] = useState<AlertHistory>();
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    useEffect(() => {

        const fetchAlertHistory = async () => {
            try {
                if (!alertInstance || alertInstance === "undefined") {
                    setError('Missing alertInstance');
                    return
                }
                const data = await getAlertHistory(alertInstance);
                console.log(data)
                setAlertHistory(data);
            } catch (e: any) {
                setError(e.message);
            } finally {
                setLoading(false);
            }
        }

        fetchAlertHistory()

    }, []);

    if (loading) return <h1>Lädt History von {alertInstance}</h1>;
    if (error) return <h1>Fehler: {error}</h1>;
    if (!alertHistory || alertHistory.alerts.length === 0) return <h1>Keine Alert History für {alertInstance} gefunden</h1>;

    return (
        <>
            <div onClick={() => navigate('/')} style={{ display: 'flex' }}>
                <button style={{ marginRight: 'auto' }}>BACK</button>
            </div>
            <h1>{formatAlertInstance(alertHistory.instance)}</h1>
            <AlertHistoryTable alertHistory={alertHistory}></AlertHistoryTable>
        </>
    );
}
