import { useEffect, useState } from 'react';
import type { Rule, RulesResponse } from '../types/Rule';
import { getRules } from '../api/GetRules';
import RuleTableItem from './RuleTableItem';
import './RuleTable.css'

interface RulesTableProps {
    reloadKey: Number;
}
export default function RulesTable({ reloadKey }: RulesTableProps) {
    const [rules, setRules] = useState<Rule[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {

        const fetchRules = async () => {
            try {
                const data: RulesResponse = await getRules()

                setRules(data.rules)
            }
            catch (err: any) {
                console.error(err)
                setError(err.message)
                setRules([])
            }
            finally {
                setLoading(false)
            }
        }

        fetchRules()

    }, [reloadKey]);

    if (loading) return <h1>Lädt aktuelle Regeln</h1>;
    if (error) return <h1>Fehler: {error}</h1>;
    if (rules === undefined || rules.length === 0) return <h1>Keine Regeln gefunden</h1>;

    return (
        <>
            <h2>Aktuelle Regeln</h2>
            <div className='rule-table-wrapper'>
                <table>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Beschreibung</th>
                            <th>Priorität</th>
                            <th>Aktiv</th>
                        </tr>
                    </thead>
                    <tbody>
                        {rules.map((rule: Rule) => (
                            <RuleTableItem key={rule.id} rule={rule} />
                        ))}
                    </tbody>
                </table>
            </div>
        </>
    );
}
