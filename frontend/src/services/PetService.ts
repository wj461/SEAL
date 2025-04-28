import { Window } from '@wailsio/runtime';

export interface IPetAction {
  update(): Promise<void>;
}

export class IdleAction implements IPetAction {
  async update(): Promise<void> {
    // 不做任何事
    return;
  }
}

export class LeftMoveAction implements IPetAction {
  private moveX: number = -10;
  private moveY: number = 0;
  private frameInterval: number = 1000 / 8;
  private prevTime: number = 0;

  async update(): Promise<void> {
    if (Date.now() - this.prevTime < this.frameInterval) {
      return;
    }
    this.prevTime = Date.now();

    const pos = await Window.Position();
    await Window.SetPosition(pos.x + this.moveX, pos.y + this.moveY);
  }
}

export class RightMoveAction implements IPetAction {
  private moveX: number = 10;
  private moveY: number = 0;
  private frameInterval: number = 1000 / 8;
  private prevTime: number = 0;

  async update(): Promise<void> {
    if (Date.now() - this.prevTime < this.frameInterval) {
      return;
    }
    this.prevTime = Date.now();

    const pos = await Window.Position();
    await Window.SetPosition(pos.x + this.moveX, pos.y + this.moveY);
  }
}

export class UpMoveAction implements IPetAction {
  private moveX: number = 0;
  private moveY: number = -10;
  private frameInterval: number = 1000 / 8;
  private prevTime: number = 0;

  async update(): Promise<void> {
    if (Date.now() - this.prevTime < this.frameInterval) {
      return;
    }
    this.prevTime = Date.now();

    const pos = await Window.Position();
    await Window.SetPosition(pos.x + this.moveX, pos.y + this.moveY);
  }
}

export class FailAction implements IPetAction {
  private moveX: number = 0;
  private y: number = 0;
  private v: number = 2000;
  private a: number = 1000;
  private dt: number = 0.016;
  private isFirst: boolean = true;
  private frameInterval: number = 1000 / 10;
  private prevTime: number = 0;

  async update(): Promise<void> {
    if (this.isFirst) {
      this.isFirst = false;
      const pos = await Window.Position();
      this.y = pos.y;
      return;
    }

    if (Date.now() - this.prevTime < this.frameInterval) {
      return;
    }
    this.prevTime = Date.now();

    this.v += this.a * this.dt;
    this.y += this.v * this.dt;

    // 模擬 time.Sleep 的效果
    await new Promise(resolve => setTimeout(resolve, this.dt * 1000));

    const pos = await Window.Position();
    await Window.SetPosition(pos.x + 2, Math.floor(this.y));
  }
}

export class PetObject {
  private actionName: string;
  private currentAction: IPetAction;

  constructor() {
    this.actionName = 'idle';
    this.currentAction = new IdleAction();
  }

  setAction(action: IPetAction) {
    this.currentAction = action;
  }

  getActionName(): string {
    return this.actionName;
  }

  getCurrentAction(): IPetAction {
    return this.currentAction;
  }
}

export class FrontendPetService {
  private actionFactory: Record<string, () => IPetAction> = {};
  private pet = new PetObject();

  constructor(action = 'failing') {
    // 註冊各種不同行為的動作
    this.actionFactory['idle'] = () => new IdleAction();
    this.actionFactory['focused_ground'] = () => new IdleAction();
    this.actionFactory['left_dragging'] = () => new IdleAction();
    this.actionFactory['right_dragging'] = () => new IdleAction();
    this.actionFactory['mid_dragging'] = () => new IdleAction();

    this.actionFactory['left_ground_walk'] = () => new LeftMoveAction();
    this.actionFactory['right_ground_walk'] = () => new RightMoveAction();
    this.actionFactory['right_up_walk'] = () => new RightMoveAction();
    this.actionFactory['failing'] = () => new FailAction();
    this.actionFactory['left_walk'] = () => new UpMoveAction();
    

    console.log("action", action);
    this.pet.setAction(this.actionFactory[action]());
  }

  setAction(action: string) {
    if (!this.pet) return;

    const actionCreator = this.actionFactory[action];
    if (!actionCreator) {
      console.error(`找不到對應的動作: ${action}`);
      return;
    }

    this.pet.setAction(actionCreator());
  }

  getState(id: number): string {
    if (!this.pet) return "pet not found";
    return this.pet.getActionName();
  }

  async update() {
      const action = this.pet.getCurrentAction();
      if (action) {
        try {
          await action.update();
        } catch (err) {
          console.error("更新時發生錯誤:", err);
        }
    }
  }
}
