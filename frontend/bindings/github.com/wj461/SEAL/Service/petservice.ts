// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as Pet$0 from "../Pet/models.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function AddPet(id: number, windowName: string): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(4040478084, id, windowName) as any;
    return $resultPromise;
}

export function GeneratePetId(): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3774493597) as any;
    return $resultPromise;
}

export function GetPetById(id: number): Promise<Pet$0.PetObject | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3391248973, id) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetScreenBounds(): Promise<$models.Bound> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3419532485) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetState(id: number): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1991151429, id) as any;
    return $resultPromise;
}

export function IdForNewWindow(): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(4241835478) as any;
    return $resultPromise;
}

export function NewPetForFrontend(windowName: string): Promise<[number, $models.Bound]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2215444106, windowName) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        $result[1] = $$createType2($result[1]);
        return $result;
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SetAction(id: number, action: string): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3241413550, id, action) as any;
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = Pet$0.PetObject.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $models.Bound.createFrom;
