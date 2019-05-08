package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyStack_Pop(t *testing.T) {
	type fields struct {
		list []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{name: "pop_empty", fields: fields{list: []int{}}, want: 0, wantErr: true},
		{name: "pop_normal_one", fields: fields{list: []int{1}}, want: 1, wantErr: false},
		{name: "pop_normal_two", fields: fields{list: []int{1, 2}}, want: 2, wantErr: false},
		{name: "pop_normal_more", fields: fields{list: []int{1, 2, 3}}, want: 3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyStack{
				list: tt.fields.list,
			}
			got, err := m.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("MyStack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MyStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Push(t *testing.T) {
	type fields struct {
		list []int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "push_empty", fields: fields{list: nil}, args: args{0},
		},
		{
			name: "push_normal_one", fields: fields{list: nil}, args: args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyStack{
				list: tt.fields.list,
			}
			m.Push(tt.args.value)
		})
	}
}

func TestMyStack_Push2(t *testing.T) {
	stack := new(MyStack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	value, err := stack.Pop()
	assert.EqualValues(t, 3, value)
	assert.NoError(t, err)
	value, err = stack.Pop()
	assert.EqualValues(t, 2, value)
	assert.NoError(t, err)
	value, err = stack.Pop()
	assert.EqualValues(t, 1, value)
	assert.NoError(t, err)
	value, err = stack.Pop()
	assert.EqualValues(t, 0, value)
	assert.Error(t, err)
}
