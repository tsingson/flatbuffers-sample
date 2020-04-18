// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MONSTER_MYGAME_EXAMPLE_H_
#define FLATBUFFERS_GENERATED_MONSTER_MYGAME_EXAMPLE_H_

#include "flatbuffers/flatbuffers.h"

#include "weapons_generated.h"

namespace Mygame {
namespace Example {

struct Vec3;

struct Monster;
struct MonsterBuilder;
struct MonsterT;

struct Weapon;
struct WeaponBuilder;
struct WeaponT;

struct SpaceShip;
struct SpaceShipBuilder;
struct SpaceShipT;

enum Color {
  Color_Red = 0,
  Color_Green = 1,
  Color_Blue = 2,
  Color_MIN = Color_Red,
  Color_MAX = Color_Blue
};

inline const Color (&EnumValuesColor())[3] {
  static const Color values[] = {
    Color_Red,
    Color_Green,
    Color_Blue
  };
  return values;
}

inline const char * const *EnumNamesColor() {
  static const char * const names[4] = {
    "Red",
    "Green",
    "Blue",
    nullptr
  };
  return names;
}

inline const char *EnumNameColor(Color e) {
  if (flatbuffers::IsOutRange(e, Color_Red, Color_Blue)) return "";
  const size_t index = static_cast<size_t>(e);
  return EnumNamesColor()[index];
}

enum Equipment {
  Equipment_NONE = 0,
  Equipment_MuLan = 1,
  Equipment_Weapon = 2,
  Equipment_Gun = 3,
  Equipment_SpaceShip = 4,
  Equipment_Other = 5,
  Equipment_MIN = Equipment_NONE,
  Equipment_MAX = Equipment_Other
};

inline const Equipment (&EnumValuesEquipment())[6] {
  static const Equipment values[] = {
    Equipment_NONE,
    Equipment_MuLan,
    Equipment_Weapon,
    Equipment_Gun,
    Equipment_SpaceShip,
    Equipment_Other
  };
  return values;
}

inline const char * const *EnumNamesEquipment() {
  static const char * const names[7] = {
    "NONE",
    "MuLan",
    "Weapon",
    "Gun",
    "SpaceShip",
    "Other",
    nullptr
  };
  return names;
}

inline const char *EnumNameEquipment(Equipment e) {
  if (flatbuffers::IsOutRange(e, Equipment_NONE, Equipment_Other)) return "";
  const size_t index = static_cast<size_t>(e);
  return EnumNamesEquipment()[index];
}

struct EquipmentUnion {
  Equipment type;
  void *value;

  EquipmentUnion() : type(Equipment_NONE), value(nullptr) {}
  EquipmentUnion(EquipmentUnion&& u) FLATBUFFERS_NOEXCEPT :
    type(Equipment_NONE), value(nullptr)
    { std::swap(type, u.type); std::swap(value, u.value); }
  EquipmentUnion(const EquipmentUnion &);
  EquipmentUnion &operator=(const EquipmentUnion &u)
    { EquipmentUnion t(u); std::swap(type, t.type); std::swap(value, t.value); return *this; }
  EquipmentUnion &operator=(EquipmentUnion &&u) FLATBUFFERS_NOEXCEPT
    { std::swap(type, u.type); std::swap(value, u.value); return *this; }
  ~EquipmentUnion() { Reset(); }

  void Reset();

  static void *UnPack(const void *obj, Equipment type, const flatbuffers::resolver_function_t *resolver);
  flatbuffers::Offset<void> Pack(flatbuffers::FlatBufferBuilder &_fbb, const flatbuffers::rehasher_function_t *_rehasher = nullptr) const;

  Mygame::Example::WeaponT *AsMuLan() {
    return type == Equipment_MuLan ?
      reinterpret_cast<Mygame::Example::WeaponT *>(value) : nullptr;
  }
  const Mygame::Example::WeaponT *AsMuLan() const {
    return type == Equipment_MuLan ?
      reinterpret_cast<const Mygame::Example::WeaponT *>(value) : nullptr;
  }
  Mygame::Example::WeaponT *AsWeapon() {
    return type == Equipment_Weapon ?
      reinterpret_cast<Mygame::Example::WeaponT *>(value) : nullptr;
  }
  const Mygame::Example::WeaponT *AsWeapon() const {
    return type == Equipment_Weapon ?
      reinterpret_cast<const Mygame::Example::WeaponT *>(value) : nullptr;
  }
  weapons::GunT *AsGun() {
    return type == Equipment_Gun ?
      reinterpret_cast<weapons::GunT *>(value) : nullptr;
  }
  const weapons::GunT *AsGun() const {
    return type == Equipment_Gun ?
      reinterpret_cast<const weapons::GunT *>(value) : nullptr;
  }
  Mygame::Example::SpaceShipT *AsSpaceShip() {
    return type == Equipment_SpaceShip ?
      reinterpret_cast<Mygame::Example::SpaceShipT *>(value) : nullptr;
  }
  const Mygame::Example::SpaceShipT *AsSpaceShip() const {
    return type == Equipment_SpaceShip ?
      reinterpret_cast<const Mygame::Example::SpaceShipT *>(value) : nullptr;
  }
  std::string *AsOther() {
    return type == Equipment_Other ?
      reinterpret_cast<std::string *>(value) : nullptr;
  }
  const std::string *AsOther() const {
    return type == Equipment_Other ?
      reinterpret_cast<const std::string *>(value) : nullptr;
  }
};

bool VerifyEquipment(flatbuffers::Verifier &verifier, const void *obj, Equipment type);
bool VerifyEquipmentVector(flatbuffers::Verifier &verifier, const flatbuffers::Vector<flatbuffers::Offset<void>> *values, const flatbuffers::Vector<uint8_t> *types);

FLATBUFFERS_MANUALLY_ALIGNED_STRUCT(4) Vec3 FLATBUFFERS_FINAL_CLASS {
 private:
  float x_;
  float y_;
  float z_;

 public:
  Vec3() {
    memset(static_cast<void *>(this), 0, sizeof(Vec3));
  }
  Vec3(float _x, float _y, float _z)
      : x_(flatbuffers::EndianScalar(_x)),
        y_(flatbuffers::EndianScalar(_y)),
        z_(flatbuffers::EndianScalar(_z)) {
  }
  float x() const {
    return flatbuffers::EndianScalar(x_);
  }
  float y() const {
    return flatbuffers::EndianScalar(y_);
  }
  float z() const {
    return flatbuffers::EndianScalar(z_);
  }
};
FLATBUFFERS_STRUCT_END(Vec3, 12);

struct MonsterT : public flatbuffers::NativeTable {
  typedef Monster TableType;
  std::unique_ptr<Mygame::Example::Vec3> pos;
  int16_t mana;
  int16_t hp;
  std::string name;
  std::vector<std::string> names;
  std::vector<uint8_t> inventory;
  Mygame::Example::Color color;
  std::vector<std::unique_ptr<Mygame::Example::WeaponT>> weapons;
  Mygame::Example::EquipmentUnion equipped;
  std::vector<Mygame::Example::EquipmentUnion> vectorOfUnions;
  std::vector<Mygame::Example::Vec3> path;
  MonsterT()
      : mana(150),
        hp(100),
        color(Mygame::Example::Color_Blue) {
  }
};

struct Monster FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef MonsterT NativeTableType;
  typedef MonsterBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_POS = 4,
    VT_MANA = 6,
    VT_HP = 8,
    VT_NAME = 10,
    VT_NAMES = 12,
    VT_INVENTORY = 16,
    VT_COLOR = 18,
    VT_WEAPONS = 20,
    VT_EQUIPPED_TYPE = 22,
    VT_EQUIPPED = 24,
    VT_VECTOROFUNIONS_TYPE = 26,
    VT_VECTOROFUNIONS = 28,
    VT_PATH = 30
  };
  const Mygame::Example::Vec3 *pos() const {
    return GetStruct<const Mygame::Example::Vec3 *>(VT_POS);
  }
  int16_t mana() const {
    return GetField<int16_t>(VT_MANA, 150);
  }
  int16_t hp() const {
    return GetField<int16_t>(VT_HP, 100);
  }
  const flatbuffers::String *name() const {
    return GetPointer<const flatbuffers::String *>(VT_NAME);
  }
  const flatbuffers::Vector<flatbuffers::Offset<flatbuffers::String>> *names() const {
    return GetPointer<const flatbuffers::Vector<flatbuffers::Offset<flatbuffers::String>> *>(VT_NAMES);
  }
  const flatbuffers::Vector<uint8_t> *inventory() const {
    return GetPointer<const flatbuffers::Vector<uint8_t> *>(VT_INVENTORY);
  }
  Mygame::Example::Color color() const {
    return static_cast<Mygame::Example::Color>(GetField<int8_t>(VT_COLOR, 2));
  }
  const flatbuffers::Vector<flatbuffers::Offset<Mygame::Example::Weapon>> *weapons() const {
    return GetPointer<const flatbuffers::Vector<flatbuffers::Offset<Mygame::Example::Weapon>> *>(VT_WEAPONS);
  }
  Mygame::Example::Equipment equipped_type() const {
    return static_cast<Mygame::Example::Equipment>(GetField<uint8_t>(VT_EQUIPPED_TYPE, 0));
  }
  const void *equipped() const {
    return GetPointer<const void *>(VT_EQUIPPED);
  }
  const Mygame::Example::Weapon *equipped_as_MuLan() const {
    return equipped_type() == Mygame::Example::Equipment_MuLan ? static_cast<const Mygame::Example::Weapon *>(equipped()) : nullptr;
  }
  const Mygame::Example::Weapon *equipped_as_Weapon() const {
    return equipped_type() == Mygame::Example::Equipment_Weapon ? static_cast<const Mygame::Example::Weapon *>(equipped()) : nullptr;
  }
  const weapons::Gun *equipped_as_Gun() const {
    return equipped_type() == Mygame::Example::Equipment_Gun ? static_cast<const weapons::Gun *>(equipped()) : nullptr;
  }
  const Mygame::Example::SpaceShip *equipped_as_SpaceShip() const {
    return equipped_type() == Mygame::Example::Equipment_SpaceShip ? static_cast<const Mygame::Example::SpaceShip *>(equipped()) : nullptr;
  }
  const flatbuffers::String *equipped_as_Other() const {
    return equipped_type() == Mygame::Example::Equipment_Other ? static_cast<const flatbuffers::String *>(equipped()) : nullptr;
  }
  const flatbuffers::Vector<uint8_t> *vectorOfUnions_type() const {
    return GetPointer<const flatbuffers::Vector<uint8_t> *>(VT_VECTOROFUNIONS_TYPE);
  }
  const flatbuffers::Vector<flatbuffers::Offset<void>> *vectorOfUnions() const {
    return GetPointer<const flatbuffers::Vector<flatbuffers::Offset<void>> *>(VT_VECTOROFUNIONS);
  }
  const flatbuffers::Vector<const Mygame::Example::Vec3 *> *path() const {
    return GetPointer<const flatbuffers::Vector<const Mygame::Example::Vec3 *> *>(VT_PATH);
  }
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyField<Mygame::Example::Vec3>(verifier, VT_POS) &&
           VerifyField<int16_t>(verifier, VT_MANA) &&
           VerifyField<int16_t>(verifier, VT_HP) &&
           VerifyOffset(verifier, VT_NAME) &&
           verifier.VerifyString(name()) &&
           VerifyOffset(verifier, VT_NAMES) &&
           verifier.VerifyVector(names()) &&
           verifier.VerifyVectorOfStrings(names()) &&
           VerifyOffset(verifier, VT_INVENTORY) &&
           verifier.VerifyVector(inventory()) &&
           VerifyField<int8_t>(verifier, VT_COLOR) &&
           VerifyOffset(verifier, VT_WEAPONS) &&
           verifier.VerifyVector(weapons()) &&
           verifier.VerifyVectorOfTables(weapons()) &&
           VerifyField<uint8_t>(verifier, VT_EQUIPPED_TYPE) &&
           VerifyOffset(verifier, VT_EQUIPPED) &&
           VerifyEquipment(verifier, equipped(), equipped_type()) &&
           VerifyOffset(verifier, VT_VECTOROFUNIONS_TYPE) &&
           verifier.VerifyVector(vectorOfUnions_type()) &&
           VerifyOffset(verifier, VT_VECTOROFUNIONS) &&
           verifier.VerifyVector(vectorOfUnions()) &&
           VerifyEquipmentVector(verifier, vectorOfUnions(), vectorOfUnions_type()) &&
           VerifyOffset(verifier, VT_PATH) &&
           verifier.VerifyVector(path()) &&
           verifier.EndTable();
  }
  MonsterT *UnPack(const flatbuffers::resolver_function_t *_resolver = nullptr) const;
  void UnPackTo(MonsterT *_o, const flatbuffers::resolver_function_t *_resolver = nullptr) const;
  static flatbuffers::Offset<Monster> Pack(flatbuffers::FlatBufferBuilder &_fbb, const MonsterT* _o, const flatbuffers::rehasher_function_t *_rehasher = nullptr);
};

struct MonsterBuilder {
  typedef Monster Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  void add_pos(const Mygame::Example::Vec3 *pos) {
    fbb_.AddStruct(Monster::VT_POS, pos);
  }
  void add_mana(int16_t mana) {
    fbb_.AddElement<int16_t>(Monster::VT_MANA, mana, 150);
  }
  void add_hp(int16_t hp) {
    fbb_.AddElement<int16_t>(Monster::VT_HP, hp, 100);
  }
  void add_name(flatbuffers::Offset<flatbuffers::String> name) {
    fbb_.AddOffset(Monster::VT_NAME, name);
  }
  void add_names(flatbuffers::Offset<flatbuffers::Vector<flatbuffers::Offset<flatbuffers::String>>> names) {
    fbb_.AddOffset(Monster::VT_NAMES, names);
  }
  void add_inventory(flatbuffers::Offset<flatbuffers::Vector<uint8_t>> inventory) {
    fbb_.AddOffset(Monster::VT_INVENTORY, inventory);
  }
  void add_color(Mygame::Example::Color color) {
    fbb_.AddElement<int8_t>(Monster::VT_COLOR, static_cast<int8_t>(color), 2);
  }
  void add_weapons(flatbuffers::Offset<flatbuffers::Vector<flatbuffers::Offset<Mygame::Example::Weapon>>> weapons) {
    fbb_.AddOffset(Monster::VT_WEAPONS, weapons);
  }
  void add_equipped_type(Mygame::Example::Equipment equipped_type) {
    fbb_.AddElement<uint8_t>(Monster::VT_EQUIPPED_TYPE, static_cast<uint8_t>(equipped_type), 0);
  }
  void add_equipped(flatbuffers::Offset<void> equipped) {
    fbb_.AddOffset(Monster::VT_EQUIPPED, equipped);
  }
  void add_vectorOfUnions_type(flatbuffers::Offset<flatbuffers::Vector<uint8_t>> vectorOfUnions_type) {
    fbb_.AddOffset(Monster::VT_VECTOROFUNIONS_TYPE, vectorOfUnions_type);
  }
  void add_vectorOfUnions(flatbuffers::Offset<flatbuffers::Vector<flatbuffers::Offset<void>>> vectorOfUnions) {
    fbb_.AddOffset(Monster::VT_VECTOROFUNIONS, vectorOfUnions);
  }
  void add_path(flatbuffers::Offset<flatbuffers::Vector<const Mygame::Example::Vec3 *>> path) {
    fbb_.AddOffset(Monster::VT_PATH, path);
  }
  explicit MonsterBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  flatbuffers::Offset<Monster> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<Monster>(end);
    return o;
  }
};

inline flatbuffers::Offset<Monster> CreateMonster(
    flatbuffers::FlatBufferBuilder &_fbb,
    const Mygame::Example::Vec3 *pos = 0,
    int16_t mana = 150,
    int16_t hp = 100,
    flatbuffers::Offset<flatbuffers::String> name = 0,
    flatbuffers::Offset<flatbuffers::Vector<flatbuffers::Offset<flatbuffers::String>>> names = 0,
    flatbuffers::Offset<flatbuffers::Vector<uint8_t>> inventory = 0,
    Mygame::Example::Color color = Mygame::Example::Color_Blue,
    flatbuffers::Offset<flatbuffers::Vector<flatbuffers::Offset<Mygame::Example::Weapon>>> weapons = 0,
    Mygame::Example::Equipment equipped_type = Mygame::Example::Equipment_NONE,
    flatbuffers::Offset<void> equipped = 0,
    flatbuffers::Offset<flatbuffers::Vector<uint8_t>> vectorOfUnions_type = 0,
    flatbuffers::Offset<flatbuffers::Vector<flatbuffers::Offset<void>>> vectorOfUnions = 0,
    flatbuffers::Offset<flatbuffers::Vector<const Mygame::Example::Vec3 *>> path = 0) {
  MonsterBuilder builder_(_fbb);
  builder_.add_path(path);
  builder_.add_vectorOfUnions(vectorOfUnions);
  builder_.add_vectorOfUnions_type(vectorOfUnions_type);
  builder_.add_equipped(equipped);
  builder_.add_weapons(weapons);
  builder_.add_inventory(inventory);
  builder_.add_names(names);
  builder_.add_name(name);
  builder_.add_pos(pos);
  builder_.add_hp(hp);
  builder_.add_mana(mana);
  builder_.add_equipped_type(equipped_type);
  builder_.add_color(color);
  return builder_.Finish();
}

inline flatbuffers::Offset<Monster> CreateMonsterDirect(
    flatbuffers::FlatBufferBuilder &_fbb,
    const Mygame::Example::Vec3 *pos = 0,
    int16_t mana = 150,
    int16_t hp = 100,
    const char *name = nullptr,
    const std::vector<flatbuffers::Offset<flatbuffers::String>> *names = nullptr,
    const std::vector<uint8_t> *inventory = nullptr,
    Mygame::Example::Color color = Mygame::Example::Color_Blue,
    const std::vector<flatbuffers::Offset<Mygame::Example::Weapon>> *weapons = nullptr,
    Mygame::Example::Equipment equipped_type = Mygame::Example::Equipment_NONE,
    flatbuffers::Offset<void> equipped = 0,
    const std::vector<uint8_t> *vectorOfUnions_type = nullptr,
    const std::vector<flatbuffers::Offset<void>> *vectorOfUnions = nullptr,
    const std::vector<Mygame::Example::Vec3> *path = nullptr) {
  auto name__ = name ? _fbb.CreateString(name) : 0;
  auto names__ = names ? _fbb.CreateVector<flatbuffers::Offset<flatbuffers::String>>(*names) : 0;
  auto inventory__ = inventory ? _fbb.CreateVector<uint8_t>(*inventory) : 0;
  auto weapons__ = weapons ? _fbb.CreateVector<flatbuffers::Offset<Mygame::Example::Weapon>>(*weapons) : 0;
  auto vectorOfUnions_type__ = vectorOfUnions_type ? _fbb.CreateVector<uint8_t>(*vectorOfUnions_type) : 0;
  auto vectorOfUnions__ = vectorOfUnions ? _fbb.CreateVector<flatbuffers::Offset<void>>(*vectorOfUnions) : 0;
  auto path__ = path ? _fbb.CreateVectorOfStructs<Mygame::Example::Vec3>(*path) : 0;
  return Mygame::Example::CreateMonster(
      _fbb,
      pos,
      mana,
      hp,
      name__,
      names__,
      inventory__,
      color,
      weapons__,
      equipped_type,
      equipped,
      vectorOfUnions_type__,
      vectorOfUnions__,
      path__);
}

flatbuffers::Offset<Monster> CreateMonster(flatbuffers::FlatBufferBuilder &_fbb, const MonsterT *_o, const flatbuffers::rehasher_function_t *_rehasher = nullptr);

struct WeaponT : public flatbuffers::NativeTable {
  typedef Weapon TableType;
  std::unique_ptr<Mygame::Example::Vec3> size;
  Mygame::Example::Color color;
  int32_t power;
  std::string name;
  WeaponT()
      : color(Mygame::Example::Color_Red),
        power(0) {
  }
};

struct Weapon FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef WeaponT NativeTableType;
  typedef WeaponBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_SIZE = 4,
    VT_COLOR = 6,
    VT_POWER = 8,
    VT_NAME = 10
  };
  const Mygame::Example::Vec3 *size() const {
    return GetStruct<const Mygame::Example::Vec3 *>(VT_SIZE);
  }
  Mygame::Example::Color color() const {
    return static_cast<Mygame::Example::Color>(GetField<int8_t>(VT_COLOR, 0));
  }
  int32_t power() const {
    return GetField<int32_t>(VT_POWER, 0);
  }
  const flatbuffers::String *name() const {
    return GetPointer<const flatbuffers::String *>(VT_NAME);
  }
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyField<Mygame::Example::Vec3>(verifier, VT_SIZE) &&
           VerifyField<int8_t>(verifier, VT_COLOR) &&
           VerifyField<int32_t>(verifier, VT_POWER) &&
           VerifyOffset(verifier, VT_NAME) &&
           verifier.VerifyString(name()) &&
           verifier.EndTable();
  }
  WeaponT *UnPack(const flatbuffers::resolver_function_t *_resolver = nullptr) const;
  void UnPackTo(WeaponT *_o, const flatbuffers::resolver_function_t *_resolver = nullptr) const;
  static flatbuffers::Offset<Weapon> Pack(flatbuffers::FlatBufferBuilder &_fbb, const WeaponT* _o, const flatbuffers::rehasher_function_t *_rehasher = nullptr);
};

struct WeaponBuilder {
  typedef Weapon Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  void add_size(const Mygame::Example::Vec3 *size) {
    fbb_.AddStruct(Weapon::VT_SIZE, size);
  }
  void add_color(Mygame::Example::Color color) {
    fbb_.AddElement<int8_t>(Weapon::VT_COLOR, static_cast<int8_t>(color), 0);
  }
  void add_power(int32_t power) {
    fbb_.AddElement<int32_t>(Weapon::VT_POWER, power, 0);
  }
  void add_name(flatbuffers::Offset<flatbuffers::String> name) {
    fbb_.AddOffset(Weapon::VT_NAME, name);
  }
  explicit WeaponBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  flatbuffers::Offset<Weapon> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<Weapon>(end);
    return o;
  }
};

inline flatbuffers::Offset<Weapon> CreateWeapon(
    flatbuffers::FlatBufferBuilder &_fbb,
    const Mygame::Example::Vec3 *size = 0,
    Mygame::Example::Color color = Mygame::Example::Color_Red,
    int32_t power = 0,
    flatbuffers::Offset<flatbuffers::String> name = 0) {
  WeaponBuilder builder_(_fbb);
  builder_.add_name(name);
  builder_.add_power(power);
  builder_.add_size(size);
  builder_.add_color(color);
  return builder_.Finish();
}

inline flatbuffers::Offset<Weapon> CreateWeaponDirect(
    flatbuffers::FlatBufferBuilder &_fbb,
    const Mygame::Example::Vec3 *size = 0,
    Mygame::Example::Color color = Mygame::Example::Color_Red,
    int32_t power = 0,
    const char *name = nullptr) {
  auto name__ = name ? _fbb.CreateString(name) : 0;
  return Mygame::Example::CreateWeapon(
      _fbb,
      size,
      color,
      power,
      name__);
}

flatbuffers::Offset<Weapon> CreateWeapon(flatbuffers::FlatBufferBuilder &_fbb, const WeaponT *_o, const flatbuffers::rehasher_function_t *_rehasher = nullptr);

struct SpaceShipT : public flatbuffers::NativeTable {
  typedef SpaceShip TableType;
  std::unique_ptr<Mygame::Example::Vec3> size;
  int32_t power;
  std::string name;
  SpaceShipT()
      : power(0) {
  }
};

struct SpaceShip FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef SpaceShipT NativeTableType;
  typedef SpaceShipBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_SIZE = 4,
    VT_POWER = 6,
    VT_NAME = 8
  };
  const Mygame::Example::Vec3 *size() const {
    return GetStruct<const Mygame::Example::Vec3 *>(VT_SIZE);
  }
  int32_t power() const {
    return GetField<int32_t>(VT_POWER, 0);
  }
  const flatbuffers::String *name() const {
    return GetPointer<const flatbuffers::String *>(VT_NAME);
  }
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyField<Mygame::Example::Vec3>(verifier, VT_SIZE) &&
           VerifyField<int32_t>(verifier, VT_POWER) &&
           VerifyOffset(verifier, VT_NAME) &&
           verifier.VerifyString(name()) &&
           verifier.EndTable();
  }
  SpaceShipT *UnPack(const flatbuffers::resolver_function_t *_resolver = nullptr) const;
  void UnPackTo(SpaceShipT *_o, const flatbuffers::resolver_function_t *_resolver = nullptr) const;
  static flatbuffers::Offset<SpaceShip> Pack(flatbuffers::FlatBufferBuilder &_fbb, const SpaceShipT* _o, const flatbuffers::rehasher_function_t *_rehasher = nullptr);
};

struct SpaceShipBuilder {
  typedef SpaceShip Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  void add_size(const Mygame::Example::Vec3 *size) {
    fbb_.AddStruct(SpaceShip::VT_SIZE, size);
  }
  void add_power(int32_t power) {
    fbb_.AddElement<int32_t>(SpaceShip::VT_POWER, power, 0);
  }
  void add_name(flatbuffers::Offset<flatbuffers::String> name) {
    fbb_.AddOffset(SpaceShip::VT_NAME, name);
  }
  explicit SpaceShipBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  flatbuffers::Offset<SpaceShip> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<SpaceShip>(end);
    return o;
  }
};

inline flatbuffers::Offset<SpaceShip> CreateSpaceShip(
    flatbuffers::FlatBufferBuilder &_fbb,
    const Mygame::Example::Vec3 *size = 0,
    int32_t power = 0,
    flatbuffers::Offset<flatbuffers::String> name = 0) {
  SpaceShipBuilder builder_(_fbb);
  builder_.add_name(name);
  builder_.add_power(power);
  builder_.add_size(size);
  return builder_.Finish();
}

inline flatbuffers::Offset<SpaceShip> CreateSpaceShipDirect(
    flatbuffers::FlatBufferBuilder &_fbb,
    const Mygame::Example::Vec3 *size = 0,
    int32_t power = 0,
    const char *name = nullptr) {
  auto name__ = name ? _fbb.CreateString(name) : 0;
  return Mygame::Example::CreateSpaceShip(
      _fbb,
      size,
      power,
      name__);
}

flatbuffers::Offset<SpaceShip> CreateSpaceShip(flatbuffers::FlatBufferBuilder &_fbb, const SpaceShipT *_o, const flatbuffers::rehasher_function_t *_rehasher = nullptr);

inline MonsterT *Monster::UnPack(const flatbuffers::resolver_function_t *_resolver) const {
  std::unique_ptr<Mygame::Example::MonsterT> _o = std::unique_ptr<Mygame::Example::MonsterT>(new MonsterT());
  UnPackTo(_o.get(), _resolver);
  return _o.release();
}

inline void Monster::UnPackTo(MonsterT *_o, const flatbuffers::resolver_function_t *_resolver) const {
  (void)_o;
  (void)_resolver;
  { auto _e = pos(); if (_e) _o->pos = std::unique_ptr<Mygame::Example::Vec3>(new Mygame::Example::Vec3(*_e)); }
  { auto _e = mana(); _o->mana = _e; }
  { auto _e = hp(); _o->hp = _e; }
  { auto _e = name(); if (_e) _o->name = _e->str(); }
  { auto _e = names(); if (_e) { _o->names.resize(_e->size()); for (flatbuffers::uoffset_t _i = 0; _i < _e->size(); _i++) { _o->names[_i] = _e->Get(_i)->str(); } } }
  { auto _e = inventory(); if (_e) { _o->inventory.resize(_e->size()); for (flatbuffers::uoffset_t _i = 0; _i < _e->size(); _i++) { _o->inventory[_i] = _e->Get(_i); } } }
  { auto _e = color(); _o->color = _e; }
  { auto _e = weapons(); if (_e) { _o->weapons.resize(_e->size()); for (flatbuffers::uoffset_t _i = 0; _i < _e->size(); _i++) { _o->weapons[_i] = std::unique_ptr<Mygame::Example::WeaponT>(_e->Get(_i)->UnPack(_resolver)); } } }
  { auto _e = equipped_type(); _o->equipped.type = _e; }
  { auto _e = equipped(); if (_e) _o->equipped.value = Mygame::Example::EquipmentUnion::UnPack(_e, equipped_type(), _resolver); }
  { auto _e = vectorOfUnions_type(); if (_e) { _o->vectorOfUnions.resize(_e->size()); for (flatbuffers::uoffset_t _i = 0; _i < _e->size(); _i++) { _o->vectorOfUnions[_i].type = static_cast<Mygame::Example::Equipment>(_e->Get(_i)); } } }
  { auto _e = vectorOfUnions(); if (_e) { _o->vectorOfUnions.resize(_e->size()); for (flatbuffers::uoffset_t _i = 0; _i < _e->size(); _i++) { _o->vectorOfUnions[_i].value = Mygame::Example::EquipmentUnion::UnPack(_e->Get(_i), vectorOfUnions_type()->GetEnum<Equipment>(_i), _resolver); } } }
  { auto _e = path(); if (_e) { _o->path.resize(_e->size()); for (flatbuffers::uoffset_t _i = 0; _i < _e->size(); _i++) { _o->path[_i] = *_e->Get(_i); } } }
}

inline flatbuffers::Offset<Monster> Monster::Pack(flatbuffers::FlatBufferBuilder &_fbb, const MonsterT* _o, const flatbuffers::rehasher_function_t *_rehasher) {
  return CreateMonster(_fbb, _o, _rehasher);
}

inline flatbuffers::Offset<Monster> CreateMonster(flatbuffers::FlatBufferBuilder &_fbb, const MonsterT *_o, const flatbuffers::rehasher_function_t *_rehasher) {
  (void)_rehasher;
  (void)_o;
  struct _VectorArgs { flatbuffers::FlatBufferBuilder *__fbb; const MonsterT* __o; const flatbuffers::rehasher_function_t *__rehasher; } _va = { &_fbb, _o, _rehasher}; (void)_va;
  auto _pos = _o->pos ? _o->pos.get() : 0;
  auto _mana = _o->mana;
  auto _hp = _o->hp;
  auto _name = _o->name.empty() ? 0 : _fbb.CreateString(_o->name);
  auto _names = _o->names.size() ? _fbb.CreateVectorOfStrings(_o->names) : 0;
  auto _inventory = _o->inventory.size() ? _fbb.CreateVector(_o->inventory) : 0;
  auto _color = _o->color;
  auto _weapons = _o->weapons.size() ? _fbb.CreateVector<flatbuffers::Offset<Mygame::Example::Weapon>> (_o->weapons.size(), [](size_t i, _VectorArgs *__va) { return CreateWeapon(*__va->__fbb, __va->__o->weapons[i].get(), __va->__rehasher); }, &_va ) : 0;
  auto _equipped_type = _o->equipped.type;
  auto _equipped = _o->equipped.Pack(_fbb);
  auto _vectorOfUnions_type = _o->vectorOfUnions.size() ? _fbb.CreateVector<uint8_t>(_o->vectorOfUnions.size(), [](size_t i, _VectorArgs *__va) { return static_cast<uint8_t>(__va->__o->vectorOfUnions[i].type); }, &_va) : 0;
  auto _vectorOfUnions = _o->vectorOfUnions.size() ? _fbb.CreateVector<flatbuffers::Offset<void>>(_o->vectorOfUnions.size(), [](size_t i, _VectorArgs *__va) { return __va->__o->vectorOfUnions[i].Pack(*__va->__fbb, __va->__rehasher); }, &_va) : 0;
  auto _path = _o->path.size() ? _fbb.CreateVectorOfStructs(_o->path) : 0;
  return Mygame::Example::CreateMonster(
      _fbb,
      _pos,
      _mana,
      _hp,
      _name,
      _names,
      _inventory,
      _color,
      _weapons,
      _equipped_type,
      _equipped,
      _vectorOfUnions_type,
      _vectorOfUnions,
      _path);
}

inline WeaponT *Weapon::UnPack(const flatbuffers::resolver_function_t *_resolver) const {
  std::unique_ptr<Mygame::Example::WeaponT> _o = std::unique_ptr<Mygame::Example::WeaponT>(new WeaponT());
  UnPackTo(_o.get(), _resolver);
  return _o.release();
}

inline void Weapon::UnPackTo(WeaponT *_o, const flatbuffers::resolver_function_t *_resolver) const {
  (void)_o;
  (void)_resolver;
  { auto _e = size(); if (_e) _o->size = std::unique_ptr<Mygame::Example::Vec3>(new Mygame::Example::Vec3(*_e)); }
  { auto _e = color(); _o->color = _e; }
  { auto _e = power(); _o->power = _e; }
  { auto _e = name(); if (_e) _o->name = _e->str(); }
}

inline flatbuffers::Offset<Weapon> Weapon::Pack(flatbuffers::FlatBufferBuilder &_fbb, const WeaponT* _o, const flatbuffers::rehasher_function_t *_rehasher) {
  return CreateWeapon(_fbb, _o, _rehasher);
}

inline flatbuffers::Offset<Weapon> CreateWeapon(flatbuffers::FlatBufferBuilder &_fbb, const WeaponT *_o, const flatbuffers::rehasher_function_t *_rehasher) {
  (void)_rehasher;
  (void)_o;
  struct _VectorArgs { flatbuffers::FlatBufferBuilder *__fbb; const WeaponT* __o; const flatbuffers::rehasher_function_t *__rehasher; } _va = { &_fbb, _o, _rehasher}; (void)_va;
  auto _size = _o->size ? _o->size.get() : 0;
  auto _color = _o->color;
  auto _power = _o->power;
  auto _name = _o->name.empty() ? 0 : _fbb.CreateString(_o->name);
  return Mygame::Example::CreateWeapon(
      _fbb,
      _size,
      _color,
      _power,
      _name);
}

inline SpaceShipT *SpaceShip::UnPack(const flatbuffers::resolver_function_t *_resolver) const {
  std::unique_ptr<Mygame::Example::SpaceShipT> _o = std::unique_ptr<Mygame::Example::SpaceShipT>(new SpaceShipT());
  UnPackTo(_o.get(), _resolver);
  return _o.release();
}

inline void SpaceShip::UnPackTo(SpaceShipT *_o, const flatbuffers::resolver_function_t *_resolver) const {
  (void)_o;
  (void)_resolver;
  { auto _e = size(); if (_e) _o->size = std::unique_ptr<Mygame::Example::Vec3>(new Mygame::Example::Vec3(*_e)); }
  { auto _e = power(); _o->power = _e; }
  { auto _e = name(); if (_e) _o->name = _e->str(); }
}

inline flatbuffers::Offset<SpaceShip> SpaceShip::Pack(flatbuffers::FlatBufferBuilder &_fbb, const SpaceShipT* _o, const flatbuffers::rehasher_function_t *_rehasher) {
  return CreateSpaceShip(_fbb, _o, _rehasher);
}

inline flatbuffers::Offset<SpaceShip> CreateSpaceShip(flatbuffers::FlatBufferBuilder &_fbb, const SpaceShipT *_o, const flatbuffers::rehasher_function_t *_rehasher) {
  (void)_rehasher;
  (void)_o;
  struct _VectorArgs { flatbuffers::FlatBufferBuilder *__fbb; const SpaceShipT* __o; const flatbuffers::rehasher_function_t *__rehasher; } _va = { &_fbb, _o, _rehasher}; (void)_va;
  auto _size = _o->size ? _o->size.get() : 0;
  auto _power = _o->power;
  auto _name = _o->name.empty() ? 0 : _fbb.CreateString(_o->name);
  return Mygame::Example::CreateSpaceShip(
      _fbb,
      _size,
      _power,
      _name);
}

inline bool VerifyEquipment(flatbuffers::Verifier &verifier, const void *obj, Equipment type) {
  switch (type) {
    case Equipment_NONE: {
      return true;
    }
    case Equipment_MuLan: {
      auto ptr = reinterpret_cast<const Mygame::Example::Weapon *>(obj);
      return verifier.VerifyTable(ptr);
    }
    case Equipment_Weapon: {
      auto ptr = reinterpret_cast<const Mygame::Example::Weapon *>(obj);
      return verifier.VerifyTable(ptr);
    }
    case Equipment_Gun: {
      auto ptr = reinterpret_cast<const weapons::Gun *>(obj);
      return verifier.VerifyTable(ptr);
    }
    case Equipment_SpaceShip: {
      auto ptr = reinterpret_cast<const Mygame::Example::SpaceShip *>(obj);
      return verifier.VerifyTable(ptr);
    }
    case Equipment_Other: {
      auto ptr = reinterpret_cast<const flatbuffers::String *>(obj);
      return verifier.VerifyString(ptr);
    }
    default: return true;
  }
}

inline bool VerifyEquipmentVector(flatbuffers::Verifier &verifier, const flatbuffers::Vector<flatbuffers::Offset<void>> *values, const flatbuffers::Vector<uint8_t> *types) {
  if (!values || !types) return !values && !types;
  if (values->size() != types->size()) return false;
  for (flatbuffers::uoffset_t i = 0; i < values->size(); ++i) {
    if (!VerifyEquipment(
        verifier,  values->Get(i), types->GetEnum<Equipment>(i))) {
      return false;
    }
  }
  return true;
}

inline void *EquipmentUnion::UnPack(const void *obj, Equipment type, const flatbuffers::resolver_function_t *resolver) {
  switch (type) {
    case Equipment_MuLan: {
      auto ptr = reinterpret_cast<const Mygame::Example::Weapon *>(obj);
      return ptr->UnPack(resolver);
    }
    case Equipment_Weapon: {
      auto ptr = reinterpret_cast<const Mygame::Example::Weapon *>(obj);
      return ptr->UnPack(resolver);
    }
    case Equipment_Gun: {
      auto ptr = reinterpret_cast<const weapons::Gun *>(obj);
      return ptr->UnPack(resolver);
    }
    case Equipment_SpaceShip: {
      auto ptr = reinterpret_cast<const Mygame::Example::SpaceShip *>(obj);
      return ptr->UnPack(resolver);
    }
    case Equipment_Other: {
      auto ptr = reinterpret_cast<const flatbuffers::String *>(obj);
      return new std::string(ptr->c_str(), ptr->size());
    }
    default: return nullptr;
  }
}

inline flatbuffers::Offset<void> EquipmentUnion::Pack(flatbuffers::FlatBufferBuilder &_fbb, const flatbuffers::rehasher_function_t *_rehasher) const {
  switch (type) {
    case Equipment_MuLan: {
      auto ptr = reinterpret_cast<const Mygame::Example::WeaponT *>(value);
      return CreateWeapon(_fbb, ptr, _rehasher).Union();
    }
    case Equipment_Weapon: {
      auto ptr = reinterpret_cast<const Mygame::Example::WeaponT *>(value);
      return CreateWeapon(_fbb, ptr, _rehasher).Union();
    }
    case Equipment_Gun: {
      auto ptr = reinterpret_cast<const weapons::GunT *>(value);
      return CreateGun(_fbb, ptr, _rehasher).Union();
    }
    case Equipment_SpaceShip: {
      auto ptr = reinterpret_cast<const Mygame::Example::SpaceShipT *>(value);
      return CreateSpaceShip(_fbb, ptr, _rehasher).Union();
    }
    case Equipment_Other: {
      auto ptr = reinterpret_cast<const std::string *>(value);
      return _fbb.CreateString(*ptr).Union();
    }
    default: return 0;
  }
}

inline EquipmentUnion::EquipmentUnion(const EquipmentUnion &u) : type(u.type), value(nullptr) {
  switch (type) {
    case Equipment_MuLan: {
      FLATBUFFERS_ASSERT(false);  // Mygame::Example::WeaponT not copyable.
      break;
    }
    case Equipment_Weapon: {
      FLATBUFFERS_ASSERT(false);  // Mygame::Example::WeaponT not copyable.
      break;
    }
    case Equipment_Gun: {
      value = new weapons::GunT(*reinterpret_cast<weapons::GunT *>(u.value));
      break;
    }
    case Equipment_SpaceShip: {
      FLATBUFFERS_ASSERT(false);  // Mygame::Example::SpaceShipT not copyable.
      break;
    }
    case Equipment_Other: {
      value = new std::string(*reinterpret_cast<std::string *>(u.value));
      break;
    }
    default:
      break;
  }
}

inline void EquipmentUnion::Reset() {
  switch (type) {
    case Equipment_MuLan: {
      auto ptr = reinterpret_cast<Mygame::Example::WeaponT *>(value);
      delete ptr;
      break;
    }
    case Equipment_Weapon: {
      auto ptr = reinterpret_cast<Mygame::Example::WeaponT *>(value);
      delete ptr;
      break;
    }
    case Equipment_Gun: {
      auto ptr = reinterpret_cast<weapons::GunT *>(value);
      delete ptr;
      break;
    }
    case Equipment_SpaceShip: {
      auto ptr = reinterpret_cast<Mygame::Example::SpaceShipT *>(value);
      delete ptr;
      break;
    }
    case Equipment_Other: {
      auto ptr = reinterpret_cast<std::string *>(value);
      delete ptr;
      break;
    }
    default: break;
  }
  value = nullptr;
  type = Equipment_NONE;
}

inline const Mygame::Example::Monster *GetMonster(const void *buf) {
  return flatbuffers::GetRoot<Mygame::Example::Monster>(buf);
}

inline const Mygame::Example::Monster *GetSizePrefixedMonster(const void *buf) {
  return flatbuffers::GetSizePrefixedRoot<Mygame::Example::Monster>(buf);
}

inline const char *MonsterIdentifier() {
  return "MNST";
}

inline bool MonsterBufferHasIdentifier(const void *buf) {
  return flatbuffers::BufferHasIdentifier(
      buf, MonsterIdentifier());
}

inline bool VerifyMonsterBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifyBuffer<Mygame::Example::Monster>(MonsterIdentifier());
}

inline bool VerifySizePrefixedMonsterBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifySizePrefixedBuffer<Mygame::Example::Monster>(MonsterIdentifier());
}

inline void FinishMonsterBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<Mygame::Example::Monster> root) {
  fbb.Finish(root, MonsterIdentifier());
}

inline void FinishSizePrefixedMonsterBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<Mygame::Example::Monster> root) {
  fbb.FinishSizePrefixed(root, MonsterIdentifier());
}

inline std::unique_ptr<Mygame::Example::MonsterT> UnPackMonster(
    const void *buf,
    const flatbuffers::resolver_function_t *res = nullptr) {
  return std::unique_ptr<Mygame::Example::MonsterT>(GetMonster(buf)->UnPack(res));
}

inline std::unique_ptr<Mygame::Example::MonsterT> UnPackSizePrefixedMonster(
    const void *buf,
    const flatbuffers::resolver_function_t *res = nullptr) {
  return std::unique_ptr<Mygame::Example::MonsterT>(GetSizePrefixedMonster(buf)->UnPack(res));
}

}  // namespace Example
}  // namespace Mygame

#endif  // FLATBUFFERS_GENERATED_MONSTER_MYGAME_EXAMPLE_H_