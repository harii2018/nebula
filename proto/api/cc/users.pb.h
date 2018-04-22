// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: users.proto

#ifndef PROTOBUF_users_2eproto__INCLUDED
#define PROTOBUF_users_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3005000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3005001 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include "types.pb.h"
#include "files.pb.h"
// @@protoc_insertion_point(includes)

namespace protobuf_users_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[2];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
void InitDefaultsUserImpl();
void InitDefaultsUser();
void InitDefaultsPushUserImpl();
void InitDefaultsPushUser();
inline void InitDefaults() {
  InitDefaultsUser();
  InitDefaultsPushUser();
}
}  // namespace protobuf_users_2eproto
namespace zproto {
class PushUser;
class PushUserDefaultTypeInternal;
extern PushUserDefaultTypeInternal _PushUser_default_instance_;
class User;
class UserDefaultTypeInternal;
extern UserDefaultTypeInternal _User_default_instance_;
}  // namespace zproto
namespace zproto {

// ===================================================================

class User : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:zproto.User) */ {
 public:
  User();
  virtual ~User();

  User(const User& from);

  inline User& operator=(const User& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  User(User&& from) noexcept
    : User() {
    *this = ::std::move(from);
  }

  inline User& operator=(User&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const User& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const User* internal_default_instance() {
    return reinterpret_cast<const User*>(
               &_User_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    0;

  void Swap(User* other);
  friend void swap(User& a, User& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline User* New() const PROTOBUF_FINAL { return New(NULL); }

  User* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(User* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string uid = 1;
  void clear_uid();
  static const int kUidFieldNumber = 1;
  const ::std::string& uid() const;
  void set_uid(const ::std::string& value);
  #if LANG_CXX11
  void set_uid(::std::string&& value);
  #endif
  void set_uid(const char* value);
  void set_uid(const char* value, size_t size);
  ::std::string* mutable_uid();
  ::std::string* release_uid();
  void set_allocated_uid(::std::string* uid);

  // int64 access_hash = 2;
  void clear_access_hash();
  static const int kAccessHashFieldNumber = 2;
  ::google::protobuf::int64 access_hash() const;
  void set_access_hash(::google::protobuf::int64 value);

  // @@protoc_insertion_point(class_scope:zproto.User)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr uid_;
  ::google::protobuf::int64 access_hash_;
  mutable int _cached_size_;
  friend struct ::protobuf_users_2eproto::TableStruct;
  friend void ::protobuf_users_2eproto::InitDefaultsUserImpl();
};
// -------------------------------------------------------------------

class PushUser : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:zproto.PushUser) */ {
 public:
  PushUser();
  virtual ~PushUser();

  PushUser(const PushUser& from);

  inline PushUser& operator=(const PushUser& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  PushUser(PushUser&& from) noexcept
    : PushUser() {
    *this = ::std::move(from);
  }

  inline PushUser& operator=(PushUser&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const PushUser& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const PushUser* internal_default_instance() {
    return reinterpret_cast<const PushUser*>(
               &_PushUser_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    1;

  void Swap(PushUser* other);
  friend void swap(PushUser& a, PushUser& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline PushUser* New() const PROTOBUF_FINAL { return New(NULL); }

  PushUser* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(PushUser* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string uid = 1;
  void clear_uid();
  static const int kUidFieldNumber = 1;
  const ::std::string& uid() const;
  void set_uid(const ::std::string& value);
  #if LANG_CXX11
  void set_uid(::std::string&& value);
  #endif
  void set_uid(const char* value);
  void set_uid(const char* value, size_t size);
  ::std::string* mutable_uid();
  ::std::string* release_uid();
  void set_allocated_uid(::std::string* uid);

  // string push_name = 3;
  void clear_push_name();
  static const int kPushNameFieldNumber = 3;
  const ::std::string& push_name() const;
  void set_push_name(const ::std::string& value);
  #if LANG_CXX11
  void set_push_name(::std::string&& value);
  #endif
  void set_push_name(const char* value);
  void set_push_name(const char* value, size_t size);
  ::std::string* mutable_push_name();
  ::std::string* release_push_name();
  void set_allocated_push_name(::std::string* push_name);

  // .zproto.Avatar push_avatar = 5;
  bool has_push_avatar() const;
  void clear_push_avatar();
  static const int kPushAvatarFieldNumber = 5;
  const ::zproto::Avatar& push_avatar() const;
  ::zproto::Avatar* release_push_avatar();
  ::zproto::Avatar* mutable_push_avatar();
  void set_allocated_push_avatar(::zproto::Avatar* push_avatar);

  // int64 access_hash = 2;
  void clear_access_hash();
  static const int kAccessHashFieldNumber = 2;
  ::google::protobuf::int64 access_hash() const;
  void set_access_hash(::google::protobuf::int64 value);

  // @@protoc_insertion_point(class_scope:zproto.PushUser)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr uid_;
  ::google::protobuf::internal::ArenaStringPtr push_name_;
  ::zproto::Avatar* push_avatar_;
  ::google::protobuf::int64 access_hash_;
  mutable int _cached_size_;
  friend struct ::protobuf_users_2eproto::TableStruct;
  friend void ::protobuf_users_2eproto::InitDefaultsPushUserImpl();
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// User

// string uid = 1;
inline void User::clear_uid() {
  uid_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& User::uid() const {
  // @@protoc_insertion_point(field_get:zproto.User.uid)
  return uid_.GetNoArena();
}
inline void User::set_uid(const ::std::string& value) {
  
  uid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:zproto.User.uid)
}
#if LANG_CXX11
inline void User::set_uid(::std::string&& value) {
  
  uid_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:zproto.User.uid)
}
#endif
inline void User::set_uid(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  uid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:zproto.User.uid)
}
inline void User::set_uid(const char* value, size_t size) {
  
  uid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:zproto.User.uid)
}
inline ::std::string* User::mutable_uid() {
  
  // @@protoc_insertion_point(field_mutable:zproto.User.uid)
  return uid_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* User::release_uid() {
  // @@protoc_insertion_point(field_release:zproto.User.uid)
  
  return uid_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void User::set_allocated_uid(::std::string* uid) {
  if (uid != NULL) {
    
  } else {
    
  }
  uid_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), uid);
  // @@protoc_insertion_point(field_set_allocated:zproto.User.uid)
}

// int64 access_hash = 2;
inline void User::clear_access_hash() {
  access_hash_ = GOOGLE_LONGLONG(0);
}
inline ::google::protobuf::int64 User::access_hash() const {
  // @@protoc_insertion_point(field_get:zproto.User.access_hash)
  return access_hash_;
}
inline void User::set_access_hash(::google::protobuf::int64 value) {
  
  access_hash_ = value;
  // @@protoc_insertion_point(field_set:zproto.User.access_hash)
}

// -------------------------------------------------------------------

// PushUser

// string uid = 1;
inline void PushUser::clear_uid() {
  uid_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& PushUser::uid() const {
  // @@protoc_insertion_point(field_get:zproto.PushUser.uid)
  return uid_.GetNoArena();
}
inline void PushUser::set_uid(const ::std::string& value) {
  
  uid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:zproto.PushUser.uid)
}
#if LANG_CXX11
inline void PushUser::set_uid(::std::string&& value) {
  
  uid_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:zproto.PushUser.uid)
}
#endif
inline void PushUser::set_uid(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  uid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:zproto.PushUser.uid)
}
inline void PushUser::set_uid(const char* value, size_t size) {
  
  uid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:zproto.PushUser.uid)
}
inline ::std::string* PushUser::mutable_uid() {
  
  // @@protoc_insertion_point(field_mutable:zproto.PushUser.uid)
  return uid_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* PushUser::release_uid() {
  // @@protoc_insertion_point(field_release:zproto.PushUser.uid)
  
  return uid_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void PushUser::set_allocated_uid(::std::string* uid) {
  if (uid != NULL) {
    
  } else {
    
  }
  uid_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), uid);
  // @@protoc_insertion_point(field_set_allocated:zproto.PushUser.uid)
}

// int64 access_hash = 2;
inline void PushUser::clear_access_hash() {
  access_hash_ = GOOGLE_LONGLONG(0);
}
inline ::google::protobuf::int64 PushUser::access_hash() const {
  // @@protoc_insertion_point(field_get:zproto.PushUser.access_hash)
  return access_hash_;
}
inline void PushUser::set_access_hash(::google::protobuf::int64 value) {
  
  access_hash_ = value;
  // @@protoc_insertion_point(field_set:zproto.PushUser.access_hash)
}

// string push_name = 3;
inline void PushUser::clear_push_name() {
  push_name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& PushUser::push_name() const {
  // @@protoc_insertion_point(field_get:zproto.PushUser.push_name)
  return push_name_.GetNoArena();
}
inline void PushUser::set_push_name(const ::std::string& value) {
  
  push_name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:zproto.PushUser.push_name)
}
#if LANG_CXX11
inline void PushUser::set_push_name(::std::string&& value) {
  
  push_name_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:zproto.PushUser.push_name)
}
#endif
inline void PushUser::set_push_name(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  push_name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:zproto.PushUser.push_name)
}
inline void PushUser::set_push_name(const char* value, size_t size) {
  
  push_name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:zproto.PushUser.push_name)
}
inline ::std::string* PushUser::mutable_push_name() {
  
  // @@protoc_insertion_point(field_mutable:zproto.PushUser.push_name)
  return push_name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* PushUser::release_push_name() {
  // @@protoc_insertion_point(field_release:zproto.PushUser.push_name)
  
  return push_name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void PushUser::set_allocated_push_name(::std::string* push_name) {
  if (push_name != NULL) {
    
  } else {
    
  }
  push_name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), push_name);
  // @@protoc_insertion_point(field_set_allocated:zproto.PushUser.push_name)
}

// .zproto.Avatar push_avatar = 5;
inline bool PushUser::has_push_avatar() const {
  return this != internal_default_instance() && push_avatar_ != NULL;
}
inline const ::zproto::Avatar& PushUser::push_avatar() const {
  const ::zproto::Avatar* p = push_avatar_;
  // @@protoc_insertion_point(field_get:zproto.PushUser.push_avatar)
  return p != NULL ? *p : *reinterpret_cast<const ::zproto::Avatar*>(
      &::zproto::_Avatar_default_instance_);
}
inline ::zproto::Avatar* PushUser::release_push_avatar() {
  // @@protoc_insertion_point(field_release:zproto.PushUser.push_avatar)
  
  ::zproto::Avatar* temp = push_avatar_;
  push_avatar_ = NULL;
  return temp;
}
inline ::zproto::Avatar* PushUser::mutable_push_avatar() {
  
  if (push_avatar_ == NULL) {
    push_avatar_ = new ::zproto::Avatar;
  }
  // @@protoc_insertion_point(field_mutable:zproto.PushUser.push_avatar)
  return push_avatar_;
}
inline void PushUser::set_allocated_push_avatar(::zproto::Avatar* push_avatar) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == NULL) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(push_avatar_);
  }
  if (push_avatar) {
    ::google::protobuf::Arena* submessage_arena = NULL;
    if (message_arena != submessage_arena) {
      push_avatar = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, push_avatar, submessage_arena);
    }
    
  } else {
    
  }
  push_avatar_ = push_avatar;
  // @@protoc_insertion_point(field_set_allocated:zproto.PushUser.push_avatar)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace zproto

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_users_2eproto__INCLUDED