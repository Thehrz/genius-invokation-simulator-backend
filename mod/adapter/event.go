package adapter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	adapter "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	converter "github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type eventAdapterLayer struct {
	implement.BaseEntityImplement
	triggerAt    enum.TriggerType
	canTriggered func(callbackContext context.CallbackContext) bool
	needClear    func() bool
	callback     func(context *context.CallbackContext)
}

func (e *eventAdapterLayer) TriggerAt() enum.TriggerType {
	return e.triggerAt
}

func (e *eventAdapterLayer) CanTriggered(context context.CallbackContext) bool {
	return e.canTriggered(context)
}

func (e *eventAdapterLayer) NeedClear() bool {
	return e.needClear()
}

func (e *eventAdapterLayer) Callback(context *context.CallbackContext) {
	e.callback(context)
}

type EventAdapter struct{}

func (e EventAdapter) Convert(source adapter.Event) (success bool, result model.Event) {
	adapterLayer := &eventAdapterLayer{
		BaseEntityImplement: implement.BaseEntityImplement{},
		triggerAt:           source.TriggerAt(),
		canTriggered:        source.TriggeredNow,
		needClear:           source.ClearNow,
		callback:            source.CallBack,
	}

	return true, adapterLayer
}

func NewEventAdapter() converter.Adapter[adapter.Event, model.Event] {
	return EventAdapter{}
}
