import type { Rule } from '../types/Rule';
import "./ActiveAlertTableItem.css"

interface RuleItemProps {
    rule: Rule;
}

export default function RuleTableItem({ rule }: RuleItemProps) {

    return (
        <tr>
            <td>{rule.name}</td>
            <td>{rule.description}</td>
            <td>{rule.priority}</td>
            <td>{rule.enabled ? "aktiv": "deaktiviert"}</td>
        </tr>
    );
}

