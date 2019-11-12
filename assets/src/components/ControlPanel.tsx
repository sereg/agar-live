import * as React from 'react';
import {Animal, Plant, Status} from '../const/Const';
import {Backward, Forward, Load, Play, Refresh, Save, Stop} from './icon/icon';

interface Props {
    status: Status
    countAnimal: number
    countPlant: number
    changes: (event: React.MouseEvent<HTMLDivElement>) => void
    restart: (event: React.MouseEvent<HTMLDivElement>) => void
    export: (event: React.MouseEvent<HTMLDivElement>) => void
    backward: (event: React.MouseEvent<HTMLDivElement>) => void
    import: (event: React.ChangeEvent<HTMLInputElement>) => void
    changeCount: (event: React.ChangeEvent<HTMLInputElement>) => void
    generate: (event: React.MouseEvent<HTMLButtonElement>) => void
}

export const ControlPanel = (props: Props) => {

    return (
        <div>
            <div className="row">
                <div className="col-sm" onClick={props.backward}>
                    <Backward width="50px"/>
                </div>
                <div className="col-sm" onClick={props.changes}>
                    {props.status === Status.stop ? (
                        <Play width="50px"/>
                    ) : (
                        <Stop width="50px"/>
                    )}
                </div>
                <div className="col-sm">
                    <Forward width={"50px"}/>
                </div>
            </div>
            <div className="row">
                <div className="col-sm" onClick={props.export}>
                    <Save  width="50px"/>
                </div>
                <div className="col-sm">
                    <Load width="50px"/>
                    <input id="file" type="file" onChange={props.import}/>
                </div>
                <div className="col-sm" onClick={props.restart}>
                    <Refresh width="50px"/>
                </div>
            </div>
            <div>
                <fieldset>
                    <legend>Generate</legend>
                    <p>
                        <label htmlFor="input">Count animal</label>
                        <input
                            value={props.countAnimal}
                            onChange={props.changeCount}
                            type="text"
                            placeholder="0"
                            name={Animal}
                        />
                    </p>
                    <p>
                        <label htmlFor="input">Count plant</label>
                        <input
                            value={props.countPlant}
                            onChange={props.changeCount}
                            type="text"
                            placeholder="0"
                            name={Plant}
                        />
                    </p>
                    <p>
                        <button type="submit" onClick={props.generate}>Generate</button>
                    </p>
                </fieldset>
            </div>
        </div>
)
}