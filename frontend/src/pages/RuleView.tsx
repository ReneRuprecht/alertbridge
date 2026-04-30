import { useState } from 'react';
import CreateRuleForm from '../components/CreateRuleForm';
import RuleTable from '../components/RuleTable';

export default function RuleView() {
    const [reloadKey, setReloadKey] = useState(0);

    const triggerReload = () => {
        setReloadKey((prev) => prev + 1)
    }

    return (
        <div>
            <RuleTable reloadKey={reloadKey} />
            <CreateRuleForm onSuccess={triggerReload} />
        </div>
    );
}
