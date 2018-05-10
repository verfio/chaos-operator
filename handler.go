package main

import (
	verfv1 "chaos-operator/pkg/apis/chaos/v1"

	log "github.com/Sirupsen/logrus"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{}) (message string)
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) (message string) {
	log.Info("TestHandler.ObjectCreated")

	mr := obj.(*verfv1.Chaos)

	log.WithFields(log.Fields{
		"namespace": mr.Spec.Namespace,
	}).Info("new chaos is scheduled")
	return mr.Spec.Namespace
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
